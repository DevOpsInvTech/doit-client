package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func (dc *DoitClient) GetVar(d *Domain, v *Var) (*Var, error) {
	newVar := &Var{}
	res, err := http.Get(dc.createAPIURL("var", v.Name, "", d.Name))
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &newVar)
	if res.StatusCode == 200 {
		return newVar, nil
	}
	return nil, errors.New(res.Status)
}

func (dc *DoitClient) GetVars(d *Domain) ([]*Var, error) {
	newVars := []*Var{}
	res, err := http.Get(dc.createAPIURL("vars", "", "", d.Name))
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &newVars)
	if res.StatusCode == 200 {
		return newVars, nil
	}
	return nil, errors.New(res.Status)
}

func (dc *DoitClient) CreateVar(d *Domain, v *Var) error {
	res, err := http.Post(dc.createAPIURL("var", v.Name, "there", d.Name), JSONMime, nil)
	if err != nil {
		return err
	}
	_, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}
	if res.StatusCode == 200 {
		return nil
	}
	return errors.New(res.Status)
}

func (dc *DoitClient) UpdateVar(d *Domain, v *Var) error {
	req, err := http.NewRequest("PUT", dc.createAPIURL("var", v.Name, v.Value, d.Name), nil)
	req.Header.Add("Content-Type", JSONMime)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	_, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}
	if res.StatusCode == 200 {
		return nil
	}
	return errors.New(res.Status)
}

func (dc *DoitClient) DeleteVar(d *Domain, v *Var) error {
	req, err := http.NewRequest("DELETE", dc.createAPIURL("var", v.Name, "", d.Name), nil)
	req.Header.Add("Content-Type", JSONMime)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	_, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}
	if res.StatusCode == 200 {
		return nil
	}
	return errors.New(res.Status)
}
