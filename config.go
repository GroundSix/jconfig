/**
 *
 * Copyright 2011 Numerotron Inc.
 * Use of this source code is governed by an MIT-style license
 * that can be found in the LICENSE file.
 * Developed at www.stathat.com by Patrick Crosby
 * Contact us on twitter with any questions:  twitter.com/stat_hat
 *
 * The jconfig package provides a simple, basic configuration file parser using JSON.
 * package jconfig
 *
 * @package jconfig
 * @author Patrick Crosby <twitter.com/stat_hat>
 */

package jconfig

import (
    "bytes"
    "encoding/json"
    "log"
    "os"
    "strings"
)

var replacements map[string]string = nil

type Config struct {
    data     map[string]interface{}
    filename string
}

func newConfig() *Config {
    result := new(Config)
    result.data  = make(map[string]interface{})
    return result
}

// Loads config information from a JSON file
func LoadConfig(filename string) *Config {
    result := newConfig()
    result.filename = filename
    err := result.parse()
    if err != nil {
        log.Fatalf("error loading config file %s: %s", filename, err)
    }
    return result
}

// Loads config information from a JSON string
func LoadConfigString(s string) *Config {
    result := newConfig()
    err := json.Unmarshal([]byte(s), &result.data)
    if err != nil {
        log.Fatalf("error parsing config string %s: %s", s, err)
    }
    return result
}

func (c *Config) StringMerge(s string) {
    next := LoadConfigString(s)
    c.merge(next.data)
}

func (c *Config) LoadMerge(filename string) {
    next := LoadConfig(filename)
    c.merge(next.data)
}

func (c *Config) merge(ndata map[string]interface{}) {
    for k, v := range ndata {
        c.data[k] = v
    }
}

func (c *Config) parse() error {
    f, err := os.Open(c.filename)
    if err != nil {
        return err
    }
    defer f.Close()
    b := new(bytes.Buffer)
    _, err = b.ReadFrom(f)
    if err != nil {
        return err
    }
    content := executeStringReplace(b.Bytes())
    err = json.Unmarshal(content, &c.data)
    if err != nil {
        return err
    }

    return nil
}

/**
 * Loops through all replacements and
 * modifies the JSON string
 *
 * @return nil
 */
func executeStringReplace(string_bytes []byte) []byte {
    json_string := string(string_bytes)
    for replace, with := range replacements {
        json_string = strings.Replace(json_string, replace, with, -1)
    }
    string_bytes = []byte(json_string)
    return string_bytes
}

/**
 * Returns a string for the config variable key
 * 
 * @param string JSON key
 *
 * @return string JSON value
 */
func (c *Config) GetString(key string) string {
    result, present := c.data[key]
    if !present {
        return ""
    }
    return result.(string)
}

/**
 * Returns an int for the config variable key
 * 
 * @param string JSON key
 *
 * @return int|float64 JSON value
 */
func (c *Config) GetInt(key string) int {
    x, ok := c.data[key]
    if !ok {
        return -1
    }
    return int(x.(float64))
}

/**
 * Returns a float64 for the config variable key
 * 
 * @param string JSON key
 *
 * @return float64 JSON value
 */
func (c *Config) GetFloat(key string) float64 {
    x, ok := c.data[key]
    if !ok {
        return -1
    }
    return x.(float64)
}

/**
 * Returns a bool for the config variable key
 * 
 * @param string JSON key
 *
 * @return bool JSON value
 */
func (c *Config) GetBool(key string) bool {
    x, ok := c.data[key]
    if !ok {
        return false
    }
    return x.(bool)
}

/**
 * Returns a array for the config variable key
 * 
 * @param string JSON key
 *
 * @return array JSON value
 */
func (c *Config) GetArray(key string) []interface{} {
    result, present := c.data[key]
    if !present {
        return []interface{}(nil)
    }
    return result.([]interface{})
}

/**
 * Returns a string map with type interface{}
 * values
 *
 * @param string JSON key
 *
 * @return map[string]interface{} JSON value
 */
func (c *Config) GetStringMap(key string) map[string]interface{} {
    result, present := c.data[key]
    if !present {
        return map[string]interface{}(nil)
    }
    return result.(map[string]interface{})
}

/**
 * Adds a value to the replacement map
 * to replace all occurrences of a string
 *
 * @param string the string target
 * @param string the replacement
 *
 * @return nil
 */
func AddStringReplace(replace string, with string) {
    if replacements == nil {
        replacements = make(map[string]string)
    }
    replacements[replace] = with
}
