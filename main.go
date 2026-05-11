package main
import ("encoding/json";"fmt";"os")
func main() {
	var v any
	dec := json.NewDecoder(os.Stdin)
	if err := dec.Decode(&v); err != nil {
		fmt.Fprintln(os.Stderr, "INVALID —", err)
		os.Exit(1)
	}
	b, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println("VALID")
	fmt.Println(string(b))
}
