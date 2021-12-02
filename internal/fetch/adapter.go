// Package fetch Copyright (c) 2021 ~ 2022 Ourea Go Framework
// Ourea Desc Dependency Injection And Simplified version DDD
package fetch

type DtoOutput interface {
	DtoTransform() (map[string]interface{}, error)
}

type Dict struct {
	DtoA map[string]interface{}
	DtoB map[string]interface{}
}

type DtoAdapter struct {
	Dto Dict
}

func (d *DtoAdapter) DtoTransform() (map[string]interface{}, error) {
	var output map[string]interface{}
	output = make(map[string]interface{}, len(d.Dto.DtoA))
	for k, v := range d.Dto.DtoA {
		if _, ok := d.Dto.DtoB[v.(string)]; !ok {
			continue
		}
		output[k] = d.Dto.DtoB[v.(string)]
	}
	return output, nil
}
