package redis

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"errors"
	"reflect"
	"time"
)

// encode data
func encode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// decode data
func decode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}

func (receiver *Client) SaveObject(key string, value interface{}, opts ...Option) error {
	if value == nil {
		return errors.New("can NOT save an nil object")
	}

	if key == "" {
		return errors.New("key can NOT be empty string")
	}

	o := getOption(opts...)
	expire := time.Duration(0)
	if o.ExpireDuration != nil {
		expire = *o.ExpireDuration
	}
	if o.Encoder == nil || *o.Encoder == EncoderGOB {
		v, err := encode(value)
		if err != nil {
			return err
		}
		return receiver.Client.Set(context.Background(), key, string(v), expire).Err()
	}

	if *o.Encoder == EncoderJSON {
		v, err := json.Marshal(value)
		if err != nil {
			return err
		}
		return receiver.Client.Set(context.Background(), key, string(v), expire).Err()
	}

	return errors.New("bad encoder")
}

func (receiver *Client) GetObject(key string, value interface{}, opts ...Option) error {
	if value == nil {
		return errors.New("can NOT save an nil object")
	}

	if key == "" {
		return errors.New("key can NOT be empty string")
	}

	if reflect.TypeOf(value).Kind() != reflect.Ptr {
		return errors.New("cant not get value of none ptr type")
	}

	c, err := receiver.Client.Get(context.Background(), key).Result()
	if err != nil {
		return err
	}

	o := getOption(opts...)
	if o.Encoder == nil || *o.Encoder == EncoderGOB {
		return decode([]byte(c), value)
	}

	if *o.Encoder == EncoderJSON {
		return json.Unmarshal([]byte(c), value)
	}

	return errors.New("bad decoder")
}
