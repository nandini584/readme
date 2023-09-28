package util
import ("fmt"
"os"
"io/ioutil")
func createFile(){
	file,err:= os.Create("readme.md")
		if err!=nil {
			fmt.Println(err)
			return
		}
	defer file.Close();

}

func generateMarkdown(title, description string) string {
	markdown:=fmt.Sprintf("# %s\n\n%s\n", title, description)
	return markdown
}

func writeFile(markdownContent string){
	err:= ioutil.WriteFile("readme.md", []byte(markdownContent),0644)
		if err != nil {
			fmt.Println(err)
			return
		}
}
