package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	_ "math"
	_ "sort"
	_ "strconv"
	_ "strings"
	_ "unicode"
)

type Cake struct {
	Name        string
	Time        string
	Ingredients string
}

type Cake1 struct {
	Name        string
	Stovetime   string
	Ingredients string
}

var jsonStr = `{"name": "Red Velvet Strawberry Cake", "time": "45 min", "ingredients": [{"ingredient_name": "Flour", "ingredient_count": "2", "ingredient_unit": "mugs" }, {"ingredient_name": "Strawberries", "ingredient_count": "7"}, {"ingredient_name": "Vanilla extract", "ingredient_count": "2.5", "ingredient_unit": "tablespoons"}]}`

func main() {
	data := []byte(jsonStr)
	u := &Cake{}
	json.Unmarshal(data, u)
	fmt.Printf("struct:\n\t%#v\n\n", u)

	data1 := []byte(`<recipes>
	<cake>
		<name>Red Velvet Strawberry Cake</name>
		<stovetime>40 min</stovetime>
		<ingredients>
			<item>
				<itemname>Flour</itemname>
				<itemcount>3</itemcount>
				<itemunit>cups</itemunit>
			</item>
			<item>
				<itemname>Vanilla extract</itemname>
				<itemcount>1.5</itemcount>
				<itemunit>tablespoons</itemunit>
			</item>
			<item>
				<itemname>Strawberries</itemname>
				<itemcount>7</itemcount>
				<itemunit></itemunit> <!-- itemunit may be empty  -->
			</item>`)
	u1 := &Cake1{}
	xml.Unmarshal(data1, u1)
	fmt.Printf("struct:\n\t%#v\n\n", u1)
}
