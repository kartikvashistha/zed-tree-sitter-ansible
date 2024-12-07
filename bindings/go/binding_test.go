package tree_sitter_ansible_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_ansible "github.com/tree-sitter/tree-sitter-ansible/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_ansible.Language())
	if language == nil {
		t.Errorf("Error loading Ansible grammar")
	}
}
