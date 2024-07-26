package main

import (
	"errors"
	"strconv"
	"strings"
)

func (rd *Reader) Read() (*Request, error) {
	requestLine, err := rd.reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	requestLine = strings.TrimSuffix(requestLine, "\r\n")
	parts := strings.Split(requestLine, " ")

	headers := map[string]string{}

	for {
		header, err := rd.reader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		header = strings.ToLower(strings.TrimSuffix(header, "\r\n"))
		if header == "" {
			break
		}

		keyValue := strings.Split(header, ":")
		key := strings.TrimSpace(keyValue[0])
		value := strings.TrimSpace(keyValue[1])

		headers[key] = value
	}

	var body []byte
	if parts[0] == MethodPost || parts[0] == MethodPut {
		if l, ok := headers["content-length"]; !ok {
			return nil, errors.New("Content-Length Not Found")
		} else {
			len, _ := strconv.Atoi(l)
			body = rd.readBody(len)
		}

	}

	return &Request{
		Method:  parts[0],
		Path:    parts[1],
		Version: parts[2],
		Headers: headers,
		Body:    string(body),
	}, nil
}

func (rd *Reader) readBody(len int) []byte {
	body := make([]byte, len)
	rd.reader.Read(body)

	return body
}

func (wt *Writer) Write(response *Response) (int, error) {
	bytes := []byte{}
	bytes = append(bytes, response.Version...)
	bytes = append(bytes, ' ')
	bytes = append(bytes, strconv.Itoa(response.Status)...)
	bytes = append(bytes, ' ')
	bytes = append(bytes, []byte(response.Reason)...)
	bytes = append(bytes, "\r\n"...)

	for key, value := range response.Headers {
		bytes = append(bytes, key...)
		bytes = append(bytes, ':')
		bytes = append(bytes, value...)
		bytes = append(bytes, "\r\n"...)
	}
	bytes = append(bytes, "\r\n"...)

	bytes = append(bytes, response.Body...)

	wt.writer.Write(bytes)
	return len(bytes), nil
}
