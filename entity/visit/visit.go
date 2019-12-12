package entity

// User is user models property
type Visit struct {
    ID        uint   `json:"id"`
    UserId    uint   `json:"id"`
    AreaCd    string `json:"area_code"`
    Date      string `json:"date"`
}
