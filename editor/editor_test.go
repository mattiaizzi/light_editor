package editor

import "testing"

func TestCreationLine(t *testing.T) {
	line := newLine(nil)
	if len(line.content) != 0 {
		t.Errorf("Invalid line length: %d", len(line.content))
	}
	if cap(line.content) != 64 {
		t.Errorf("Ivalid line capacity: %d", cap(line.content))
	}

	content, length, capacity := getContentMetadata("Hello Test")
	line = newLine(content)
	if len(line.content) != length {
		t.Errorf("Invalid line length: %d", len(line.content))
	}
	if cap(line.content) != capacity {
		t.Errorf("Ivalid line capacity: %d", cap(line.content))
	}

	content, length, capacity = getContentMetadata("zrtkhjywevzlejhrtzakllmfipjfjhsgzzxtxkxwuflsshxmwvwyivmbvxrqmrgchvoxojquoipsapxaslaummzncrvtj")
	line = newLine(content)
	if len(line.content) != length {
		t.Errorf("Invalid line length: %d", len(line.content))
	}
	if cap(line.content) != capacity {
		t.Errorf("Ivalid line capacity: %d", cap(line.content))
	}
}

func TestAddCharacter(t *testing.T) {
	line := newLine(nil)
	for i, char := range []byte("Hello World") {
		line.insertChar(uint(i), char)
	}
	if string(line.content) != "Hello World" {
		t.Errorf("Expected: %v. Found: %v", "Hello World", string(line.content))
	}
	line = newLine([]byte("Hell World"))
	line.insertChar(4, 'o')
	if string(line.content) != "Hello World" {
		t.Errorf("Expected: %v. Found: %v", "Hello World", string(line.content))
	}

	line = newLine([]byte("Hello Worl"))
	line.insertChar(123, 'd')
	if string(line.content) != "Hello World" {
		t.Errorf("Expected: %v. Found: %v", "Hello World", string(line.content))
	}
}

func getContentMetadata(s string) (content []byte, length int, capacity int) {
	return []byte(s), len(s), len(s) + 64 - (len(s) % 64)
}
