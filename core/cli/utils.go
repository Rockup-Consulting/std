package cli

import (
	"fmt"
	"strings"
	"text/tabwriter"
)

func (m *Menu) executePrint() error {
	if m == nil || len(m.commands) == 0 {
		fmt.Fprintf(m.this.w, "no commands in this menu \n")
		return nil
	}

	var sb strings.Builder
	tw := tabwriter.NewWriter(&sb, 30, 4, 0, ' ', tabwriter.TabIndent)

	fmt.Fprintf(&sb, "%s: %s\n", m.this.arg, m.this.d)

	// print default group - could be nil if not initialised
	if m.defaultGroup != nil {
		fmt.Fprintln(tw)
		for _, cmnd := range m.defaultGroup.commands {
			cmnd.e.Print(tw)
		}
	}

	// print groups
	for _, group := range m.groups {
		fmt.Fprintln(tw)
		if group.name != "" {
			fmt.Fprintf(tw, "%s\n", group.name)
		}

		for _, cmnd := range group.commands {
			cmnd.e.Print(tw)
		}
	}

	tw.Flush()
	fmt.Print(sb.String())

	return nil
}
