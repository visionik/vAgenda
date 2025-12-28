package builder

import "github.com/visionik/vContext/api/go/pkg/core"

// PlanBuilder provides a fluent API for building Plan documents.
type PlanBuilder struct {
	doc *core.Document
}

// NewPlan creates a new Plan builder with the specified title and version.
func NewPlan(title, version string) *PlanBuilder {
	return NewPlanWithStatus(version, title, core.PlanStatusDraft)
}

// NewPlanWithStatus creates a new Plan builder with explicit status.
//
// This matches the intent of the original extension proposal (version, title, status).
func NewPlanWithStatus(version, title string, status core.PlanStatus) *PlanBuilder {
	return &PlanBuilder{
		doc: &core.Document{
			Info: core.Info{
				Version: version,
			},
			Plan: &core.Plan{
				Title:      title,
				Status:     status,
				Narratives: make(map[string]core.Narrative),
			},
		},
	}
}

// WithAuthor sets the document author.
func (b *PlanBuilder) WithAuthor(author string) *PlanBuilder {
	b.doc.Info.Author = author
	return b
}

// WithDescription sets the document description.
func (b *PlanBuilder) WithDescription(description string) *PlanBuilder {
	b.doc.Info.Description = description
	return b
}

// WithStatus sets the plan status.
func (b *PlanBuilder) WithStatus(status core.PlanStatus) *PlanBuilder {
	b.doc.Plan.Status = status
	return b
}

// WithNarrative adds a narrative to the plan.
func (b *PlanBuilder) WithNarrative(key, title, content string) *PlanBuilder {
	b.doc.Plan.Narratives[key] = core.Narrative{
		Title:   title,
		Content: content,
	}
	return b
}

// WithProposal adds the required proposal narrative.
func (b *PlanBuilder) WithProposal(title, content string) *PlanBuilder {
	return b.WithNarrative("proposal", title, content)
}

// WithProblem adds a problem statement narrative.
func (b *PlanBuilder) WithProblem(title, content string) *PlanBuilder {
	return b.WithNarrative("problem", title, content)
}

// WithContext adds a context narrative.
func (b *PlanBuilder) WithContext(title, content string) *PlanBuilder {
	return b.WithNarrative("context", title, content)
}

// WithAlternatives adds an alternatives narrative.
func (b *PlanBuilder) WithAlternatives(title, content string) *PlanBuilder {
	return b.WithNarrative("alternatives", title, content)
}

// WithRisks adds a risks narrative.
func (b *PlanBuilder) WithRisks(title, content string) *PlanBuilder {
	return b.WithNarrative("risks", title, content)
}

// WithTesting adds a testing narrative.
func (b *PlanBuilder) WithTesting(title, content string) *PlanBuilder {
	return b.WithNarrative("testing", title, content)
}

// AddPlanItem adds a plan item to the plan.
func (b *PlanBuilder) AddPlanItem(title string, status core.PlanItemStatus) *PlanBuilder {
	item := core.PlanItem{
		Title:  title,
		Status: status,
	}
	b.doc.Plan.Items = append(b.doc.Plan.Items, item)
	return b
}

// AddPendingItem adds a pending plan item to the plan.
func (b *PlanBuilder) AddPendingItem(title string) *PlanBuilder {
	return b.AddPlanItem(title, core.PlanItemStatusPending)
}

// AddInProgressItem adds an in-progress plan item to the plan.
func (b *PlanBuilder) AddInProgressItem(title string) *PlanBuilder {
	return b.AddPlanItem(title, core.PlanItemStatusInProgress)
}

// AddCompletedItem adds a completed plan item to the plan.
func (b *PlanBuilder) AddCompletedItem(title string) *PlanBuilder {
	return b.AddPlanItem(title, core.PlanItemStatusCompleted)
}

// Build returns the constructed document.
func (b *PlanBuilder) Build() *core.Document {
	return b.doc
}
