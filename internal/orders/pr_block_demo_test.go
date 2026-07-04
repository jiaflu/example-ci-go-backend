package orders

import "testing"

func TestPRBlockedByFailingUnitTest(t *testing.T) {
	t.Fatal("intentional failure: this PR demonstrates CI blocking a merge when unit tests fail")
}
