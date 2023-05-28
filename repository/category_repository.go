package repository

import "go-unittest/entity"

/** 
kontrak interface = interface ini pny kontrak func yaitu FindById dg 
param id dan return pointer Category dari pkg entity

tujuan ada kontrak interface = 
agar testing engga hrs terhubung DB (3rd party) secara real
*/

type CategoryRepository interface {
  FindById(id string) *entity.Category
}