package main

func main(){
  cli := &CommandLine{}
	err := cli.Usage()
	if err != nil {
		t.Log("ERR")
		t.Fatal(err)
	}
	t.Log("Fin")
}
