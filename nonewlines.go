package nonewlines

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
)

type visitFn func(node ast.Node) ast.Visitor

func (fn visitFn) Visit(node ast.Node) ast.Visitor {
	return fn(node)
}

func Process(filename string, src []byte) ([]byte, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, src, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	// Create an ast.CommentMap from the ast.File's comments.
	// This helps keeping the association between comments
	// and AST nodes.
	cmap := ast.NewCommentMap(fset, f, f.Comments)

	var visitor visitFn
	visitor = visitFn(func(node ast.Node) ast.Visitor {
		if node == nil {
			return visitor
		}

		switch v := node.(type) {
		case *ast.FuncDecl:
			if len(v.Body.List) == 0 {
				return visitor
			}

			// Ignore multiline function declarations.
			if isMultilineFuncDecl(v, fset) {
				return visitor
			}

			// Trim newlines between opening brace and first statement.
			stmtLine := firstLine(v.Body.List[0], fset, cmap)
			braceLine := fset.Position(v.Body.Lbrace).Line
			numSeparatorLines := stmtLine - braceLine - 1

			for i := 0; i < numSeparatorLines; i++ {
				fset.File(v.Pos()).MergeLine(braceLine)
			}

			// Trim newlines before the closing brace.
			stmt := v.Body.List[len(v.Body.List)-1]
			stmtLine = fset.Position(stmt.End()).Line
			braceLine = fset.Position(v.Body.Rbrace).Line
			lastComment := lastCommentBetween(stmtLine, braceLine, fset, f.Comments)
			if lastComment > 0 && lastComment > stmtLine {
				stmtLine = lastComment
			}
			numSeparatorLines = braceLine - stmtLine - 1

			for i := 0; i < numSeparatorLines; i++ {
				fset.File(v.Pos()).MergeLine(stmtLine)
			}
		}

		return visitor
	})
	ast.Walk(visitor, f)

	var buf bytes.Buffer
	if err := format.Node(&buf, fset, f); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func isMultilineFuncDecl(decl *ast.FuncDecl, fset *token.FileSet) bool {
	return fset.Position(decl.Pos()).Line != fset.Position(decl.Body.Lbrace).Line
}

func firstLine(stmt ast.Stmt, fset *token.FileSet, cmap ast.CommentMap) int {
	line := fset.Position(stmt.Pos()).Line

	// Get comments associated with the statement.
	commentGroups, found := cmap[stmt]
	if !found {
		return line
	}

	for _, cg := range commentGroups {
		commentLine := fset.Position(cg.Pos()).Line
		if commentLine < line {
			line = commentLine
		}
	}

	return line
}

func lastCommentBetween(m, n int, fset *token.FileSet, comments []*ast.CommentGroup) int {
	lastLine := -1

	for _, cg := range comments {
		line := fset.Position(cg.End()).Line

		if line > m && line < n && line > lastLine {
			lastLine = line
		}

		if line > n {
			break
		}
	}

	return lastLine
}
