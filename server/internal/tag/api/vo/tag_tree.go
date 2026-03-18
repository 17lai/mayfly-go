package vo

import (
	"mayfly-go/internal/tag/application/dto"
	"strings"
)

type TagTreeVOS []*dto.SimpleTagTree

type TagTreeItem struct {
	*dto.SimpleTagTree
	Children []*TagTreeItem `json:"children"`
	NamePath string         `json:"namePath"`
}

func (m *TagTreeVOS) ToTrees(pid uint64) []*TagTreeItem {
	var ttis []*TagTreeItem
	if len(*m) == 0 {
		return ttis
	}

	tagMap := make(map[string]*TagTreeItem)
	var roots []*TagTreeItem
	for _, tag := range *m {
		tti := &TagTreeItem{SimpleTagTree: tag}
		tagMap[tag.CodePath] = tti
		ttis = append(ttis, tti)
		if tti.IsRoot() {
			roots = append(roots, tti)
			tti.Root = true
		}
	}

	for _, node := range ttis {
		if node.Root {
			continue
		}
		parentCodePath := node.GetParentPath()
		parentNode := tagMap[parentCodePath]
		if parentNode != nil {
			parentNode.Children = append(parentNode.Children, node)
		}
	}

	return roots
}

func (m *TagTreeVOS) ToFlattenTrees(pid uint64) []*TagTreeItem {
	trees := m.ToTrees(pid)
	result := make([]*TagTreeItem, 0)

	var flatten func(node *TagTreeItem, namePath []string)
	flatten = func(node *TagTreeItem, namePath []string) {
		currentNamePath := append(namePath, node.Name)

		if node.Type != -1 {
			return
		}

		hasNonMinus1Child := false
		for _, child := range node.Children {
			if child.Type != -1 {
				hasNonMinus1Child = true
				break
			}
		}

		if hasNonMinus1Child {
			newNode := &TagTreeItem{
				SimpleTagTree: node.SimpleTagTree,
				Children:      make([]*TagTreeItem, 0),
				NamePath:      strings.Join(currentNamePath, "/"),
			}
			for _, child := range node.Children {
				if child.Type != -1 {
					childCopy := &TagTreeItem{
						SimpleTagTree: child.SimpleTagTree,
						Children:      make([]*TagTreeItem, 0),
						NamePath:      strings.Join(append(currentNamePath, child.Name), "/"),
					}
					for _, grandchild := range child.Children {
						grandchildCopy := &TagTreeItem{
							SimpleTagTree: grandchild.SimpleTagTree,
							NamePath:      strings.Join(append(currentNamePath, child.Name, grandchild.Name), "/"),
						}
						childCopy.Children = append(childCopy.Children, grandchildCopy)
					}
					newNode.Children = append(newNode.Children, childCopy)
				} else {
					flatten(child, currentNamePath)
				}
			}
			result = append(result, newNode)
			return
		}

		for _, child := range node.Children {
			flatten(child, currentNamePath)
		}
	}

	for _, tree := range trees {
		flatten(tree, []string{})
	}

	return result
}
