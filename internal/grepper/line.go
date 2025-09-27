package grepper

/*
func (g *Grepper) ProcessKeys2(reader *bufio.Reader) ([]LineResult, error) {
	lineNum := 0
	var res []LineResult
	var lineResult LineResult
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, fmt.Errorf("read string error: %w", err)
		}
		lineNum++
		if g.isMatch(line) {
			lineResult = LineResult{
				text: line,
				num:  lineNum,
			}
			res = append(res, lineResult)
		}
	}
	return res, nil
}

/*
func (g *Grepper) ProcessKeys(file string) ([]LineResult, error) {
	reader := bufio.NewReader(file)


}
*/
