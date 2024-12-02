package tree_sitter_ansible_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-ansible"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_ansible.Language())
	if language == nil {
		t.Errorf("Error loading Ansible grammar")
	}
}
