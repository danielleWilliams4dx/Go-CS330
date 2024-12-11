package rpg
import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
)

//in reference to code from https://www.educative.io/answers/how-to-read-a-csv-file-in-the-go-language
func GetRecords(filePath string) [][]string{
	file, error := os.Open(filePath)
	//if there's an error, print it
	PrintError(error)
	//close the file to prevent a memory leak
	defer file.Close()

	fr := csv.NewReader(file)
	//read the file content and store it in records
	records, error := fr.ReadAll()
	PrintError(error)
	return records
} 

//prints error if there is one
func PrintError(e error){
	if(e != nil){
		fmt.Println(e)
	}
}

//adds line breaks to imported descriptions
func addLineBreak(desc string) string{
	re := regexp.MustCompile(`\. `)
	return "\t" + re.ReplaceAllString(desc, ".\n\t")
}