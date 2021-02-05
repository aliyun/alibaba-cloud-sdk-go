package facebody

type DataInGetFaceEntity struct {
	EntityId string          `json:"EntityId" xml:"EntityId"`
	DbName   string          `json:"DbName" xml:"DbName"`
	Labels   string          `json:"Labels" xml:"Labels"`
	Faces    []FaceItemsItem `json:"Faces" xml:"Faces"`
}

