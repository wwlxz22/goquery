package goquery

import (
	"testing"
)

func TestFind(t *testing.T) {
	sel := Doc().Root.Find("div.row-fluid")
	AssertLength(t, sel.Nodes, 9)
}

func TestFindRollback(t *testing.T) {
	sel := Doc().Root.Find("div.row-fluid")
	sel2 := sel.Find("a").End()
	AssertEqual(t, sel, sel2)
}

func TestFindNotSelf(t *testing.T) {
	sel := Doc().Root.Find("h1").Find("h1")
	AssertLength(t, sel.Nodes, 0)
}

func TestFindInvalidSelector(t *testing.T) {
	defer AssertPanic(t)
	Doc().Root.Find(":+ ^")
}

func TestChainedFind(t *testing.T) {
	sel := Doc().Root.Find("div.hero-unit").Find(".row-fluid")
	AssertLength(t, sel.Nodes, 4)
}

func TestChildren(t *testing.T) {
	sel := Doc().Root.Find(".pvk-content").Children()
	AssertLength(t, sel.Nodes, 5)
}

func TestChildrenRollback(t *testing.T) {
	sel := Doc().Root.Find(".pvk-content")
	sel2 := sel.Children().End()
	AssertEqual(t, sel, sel2)
}

func TestContents(t *testing.T) {
	sel := Doc().Root.Find(".pvk-content").Contents()
	AssertLength(t, sel.Nodes, 13)
}

func TestContentsRollback(t *testing.T) {
	sel := Doc().Root.Find(".pvk-content")
	sel2 := sel.Contents().End()
	AssertEqual(t, sel, sel2)
}

func TestChildrenFiltered(t *testing.T) {
	sel := Doc().Root.Find(".pvk-content").ChildrenFiltered(".hero-unit")
	AssertLength(t, sel.Nodes, 1)
}

func TestChildrenFilteredRollback(t *testing.T) {
	sel := Doc().Root.Find(".pvk-content")
	sel2 := sel.ChildrenFiltered(".hero-unit").End()
	AssertEqual(t, sel, sel2)
}

func TestContentsFiltered(t *testing.T) {
	sel := Doc().Root.Find(".pvk-content").ContentsFiltered(".hero-unit")
	AssertLength(t, sel.Nodes, 1)
}

func TestContentsFilteredRollback(t *testing.T) {
	sel := Doc().Root.Find(".pvk-content")
	sel2 := sel.ContentsFiltered(".hero-unit").End()
	AssertEqual(t, sel, sel2)
}

func TestChildrenFilteredNone(t *testing.T) {
	sel := Doc().Root.Find(".pvk-content").ChildrenFiltered("a.btn")
	AssertLength(t, sel.Nodes, 0)
}

func TestParent(t *testing.T) {
	sel := Doc().Root.Find(".container-fluid").Parent()
	AssertLength(t, sel.Nodes, 3)
}

func TestParentRollback(t *testing.T) {
	sel := Doc().Root.Find(".container-fluid")
	sel2 := sel.Parent().End()
	AssertEqual(t, sel, sel2)
}

func TestParentBody(t *testing.T) {
	sel := Doc().Root.Find("body").Parent()
	AssertLength(t, sel.Nodes, 1)
}

func TestParentFiltered(t *testing.T) {
	sel := Doc().Root.Find(".container-fluid").ParentFiltered(".hero-unit")
	AssertLength(t, sel.Nodes, 1)
	AssertClass(t, sel, "hero-unit")
}

func TestParentFilteredRollback(t *testing.T) {
	sel := Doc().Root.Find(".container-fluid")
	sel2 := sel.ParentFiltered(".hero-unit").End()
	AssertEqual(t, sel, sel2)
}

func TestParents(t *testing.T) {
	sel := Doc().Root.Find(".container-fluid").Parents()
	AssertLength(t, sel.Nodes, 8)
}

func TestParentsOrder(t *testing.T) {
	sel := Doc().Root.Find("#cf2").Parents()
	AssertLength(t, sel.Nodes, 6)
	if !sel.Eq(0).HasClass("hero-unit") {
		t.Errorf("Element at 0 should be hero-unit, found %+v.", sel.Get(0))
	}
	if !sel.Eq(1).HasClass("pvk-content") {
		t.Errorf("Element at 1 should be pvk-content, found %+v.", sel.Get(1))
	}
	if !sel.Eq(2).Is("div.row-fluid") {
		t.Errorf("Element at 2 should be row-fluid, found %+v.", sel.Get(2))
	}
	if !sel.Eq(3).Is("#cf1") {
		t.Errorf("Element at 3 should be cf1, found %+v.", sel.Get(3))
	}
	if !sel.Eq(4).Is("body") {
		t.Errorf("Element at 4 should be body, found %+v.", sel.Get(4))
	}
	if !sel.Eq(5).Is("html") {
		t.Errorf("Element at 5 should be html, found %+v.", sel.Get(5))
	}
}

func TestParentsRollback(t *testing.T) {
	sel := Doc().Root.Find(".container-fluid")
	sel2 := sel.Parents().End()
	AssertEqual(t, sel, sel2)
}

func TestParentsFiltered(t *testing.T) {
	sel := Doc().Root.Find(".container-fluid").ParentsFiltered("body")
	AssertLength(t, sel.Nodes, 1)
}

func TestParentsFilteredRollback(t *testing.T) {
	sel := Doc().Root.Find(".container-fluid")
	sel2 := sel.ParentsFiltered("body").End()
	AssertEqual(t, sel, sel2)
}

func TestParentsUntil(t *testing.T) {
	sel := Doc().Root.Find(".container-fluid").ParentsUntil("body")
	AssertLength(t, sel.Nodes, 6)
}

func TestParentsUntilRollback(t *testing.T) {
	sel := Doc().Root.Find(".container-fluid")
	sel2 := sel.ParentsUntil("body").End()
	AssertEqual(t, sel, sel2)
}

func TestParentsUntilSelection(t *testing.T) {
	sel := Doc().Root.Find(".container-fluid")
	sel2 := Doc().Root.Find(".pvk-content")
	sel = sel.ParentsUntilSelection(sel2)
	AssertLength(t, sel.Nodes, 3)
}

func TestParentsUntilSelectionRollback(t *testing.T) {
	sel := Doc().Root.Find(".container-fluid")
	sel2 := Doc().Root.Find(".pvk-content")
	sel2 = sel.ParentsUntilSelection(sel2).End()
	AssertEqual(t, sel, sel2)
}

func TestParentsUntilNodes(t *testing.T) {
	sel := Doc().Root.Find(".container-fluid")
	sel2 := Doc().Root.Find(".pvk-content, .hero-unit")
	sel = sel.ParentsUntilNodes(sel2.Nodes...)
	AssertLength(t, sel.Nodes, 2)
}

func TestParentsUntilNodesRollback(t *testing.T) {
	sel := Doc().Root.Find(".container-fluid")
	sel2 := Doc().Root.Find(".pvk-content, .hero-unit")
	sel2 = sel.ParentsUntilNodes(sel2.Nodes...).End()
	AssertEqual(t, sel, sel2)
}

func TestParentsFilteredUntil(t *testing.T) {
	sel := Doc().Root.Find(".container-fluid").ParentsFilteredUntil(".pvk-content", "body")
	AssertLength(t, sel.Nodes, 2)
}

func TestParentsFilteredUntilRollback(t *testing.T) {
	sel := Doc().Root.Find(".container-fluid")
	sel2 := sel.ParentsFilteredUntil(".pvk-content", "body").End()
	AssertEqual(t, sel, sel2)
}

func TestParentsFilteredUntilSelection(t *testing.T) {
	sel := Doc().Root.Find(".container-fluid")
	sel2 := Doc().Root.Find(".row-fluid")
	sel = sel.ParentsFilteredUntilSelection("div", sel2)
	AssertLength(t, sel.Nodes, 3)
}

func TestParentsFilteredUntilSelectionRollback(t *testing.T) {
	sel := Doc().Root.Find(".container-fluid")
	sel2 := Doc().Root.Find(".row-fluid")
	sel2 = sel.ParentsFilteredUntilSelection("div", sel2).End()
	AssertEqual(t, sel, sel2)
}

func TestParentsFilteredUntilNodes(t *testing.T) {
	sel := Doc().Root.Find(".container-fluid")
	sel2 := Doc().Root.Find(".row-fluid")
	sel = sel.ParentsFilteredUntilNodes("body", sel2.Nodes...)
	AssertLength(t, sel.Nodes, 1)
}

func TestParentsFilteredUntilNodesRollback(t *testing.T) {
	sel := Doc().Root.Find(".container-fluid")
	sel2 := Doc().Root.Find(".row-fluid")
	sel2 = sel.ParentsFilteredUntilNodes("body", sel2.Nodes...).End()
	AssertEqual(t, sel, sel2)
}

func TestSiblings(t *testing.T) {
	sel := Doc().Root.Find("h1").Siblings()
	AssertLength(t, sel.Nodes, 1)
}

func TestSiblingsRollback(t *testing.T) {
	sel := Doc().Root.Find("h1")
	sel2 := sel.Siblings().End()
	AssertEqual(t, sel, sel2)
}

func TestSiblings2(t *testing.T) {
	sel := Doc().Root.Find(".pvk-gutter").Siblings()
	AssertLength(t, sel.Nodes, 9)
}

func TestSiblings3(t *testing.T) {
	sel := Doc().Root.Find("body>.container-fluid").Siblings()
	AssertLength(t, sel.Nodes, 0)
}

func TestSiblingsFiltered(t *testing.T) {
	sel := Doc().Root.Find(".pvk-gutter").SiblingsFiltered(".pvk-content")
	AssertLength(t, sel.Nodes, 3)
}

func TestSiblingsFilteredRollback(t *testing.T) {
	sel := Doc().Root.Find(".pvk-gutter")
	sel2 := sel.SiblingsFiltered(".pvk-content").End()
	AssertEqual(t, sel, sel2)
}

func TestNext(t *testing.T) {
	sel := Doc().Root.Find("h1").Next()
	AssertLength(t, sel.Nodes, 1)
}

func TestNextRollback(t *testing.T) {
	sel := Doc().Root.Find("h1")
	sel2 := sel.Next().End()
	AssertEqual(t, sel, sel2)
}

func TestNext2(t *testing.T) {
	sel := Doc().Root.Find(".close").Next()
	AssertLength(t, sel.Nodes, 1)
}

func TestNextNone(t *testing.T) {
	sel := Doc().Root.Find("small").Next()
	AssertLength(t, sel.Nodes, 0)
}

func TestNextFiltered(t *testing.T) {
	sel := Doc().Root.Find(".container-fluid").NextFiltered("div")
	AssertLength(t, sel.Nodes, 2)
}

func TestNextFilteredRollback(t *testing.T) {
	sel := Doc().Root.Find(".container-fluid")
	sel2 := sel.NextFiltered("div").End()
	AssertEqual(t, sel, sel2)
}

func TestNextFiltered2(t *testing.T) {
	sel := Doc().Root.Find(".container-fluid").NextFiltered("[ng-view]")
	AssertLength(t, sel.Nodes, 1)
}

func TestPrev(t *testing.T) {
	sel := Doc().Root.Find(".red").Prev()
	AssertLength(t, sel.Nodes, 1)
	AssertClass(t, sel, "green")
}

func TestPrevRollback(t *testing.T) {
	sel := Doc().Root.Find(".red")
	sel2 := sel.Prev().End()
	AssertEqual(t, sel, sel2)
}

func TestPrev2(t *testing.T) {
	sel := Doc().Root.Find(".row-fluid").Prev()
	AssertLength(t, sel.Nodes, 5)
}

func TestPrevNone(t *testing.T) {
	sel := Doc().Root.Find("h2").Prev()
	AssertLength(t, sel.Nodes, 0)
}

func TestPrevFiltered(t *testing.T) {
	sel := Doc().Root.Find(".row-fluid").PrevFiltered(".row-fluid")
	AssertLength(t, sel.Nodes, 5)
}

func TestPrevFilteredRollback(t *testing.T) {
	sel := Doc().Root.Find(".row-fluid")
	sel2 := sel.PrevFiltered(".row-fluid").End()
	AssertEqual(t, sel, sel2)
}

func TestNextAll(t *testing.T) {
	sel := Doc().Root.Find("#cf2 div:nth-child(1)").NextAll()
	AssertLength(t, sel.Nodes, 3)
}

func TestNextAllRollback(t *testing.T) {
	sel := Doc().Root.Find("#cf2 div:nth-child(1)")
	sel2 := sel.NextAll().End()
	AssertEqual(t, sel, sel2)
}

func TestNextAll2(t *testing.T) {
	sel := Doc().Root.Find("div[ng-cloak]").NextAll()
	AssertLength(t, sel.Nodes, 1)
}

func TestNextAllNone(t *testing.T) {
	sel := Doc().Root.Find(".footer").NextAll()
	AssertLength(t, sel.Nodes, 0)
}

func TestNextAllFiltered(t *testing.T) {
	sel := Doc().Root.Find("#cf2 .row-fluid").NextAllFiltered("[ng-cloak]")
	AssertLength(t, sel.Nodes, 2)
}

func TestNextAllFilteredRollback(t *testing.T) {
	sel := Doc().Root.Find("#cf2 .row-fluid")
	sel2 := sel.NextAllFiltered("[ng-cloak]").End()
	AssertEqual(t, sel, sel2)
}

func TestNextAllFiltered2(t *testing.T) {
	sel := Doc().Root.Find(".close").NextAllFiltered("h4")
	AssertLength(t, sel.Nodes, 1)
}

func TestPrevAll(t *testing.T) {
	sel := Doc().Root.Find("[ng-view]").PrevAll()
	AssertLength(t, sel.Nodes, 2)
}

func TestPrevAllOrder(t *testing.T) {
	sel := Doc().Root.Find("[ng-view]").PrevAll()
	AssertLength(t, sel.Nodes, 2)
	if !sel.Eq(0).Is("#cf4") {
		t.Errorf("Element at 0 should be cf4, found %+v.", sel.Get(0))
	}
	if !sel.Eq(1).Is("#cf3") {
		t.Errorf("Element at 1 should be cf3, found %+v.", sel.Get(1))
	}
}

func TestPrevAllRollback(t *testing.T) {
	sel := Doc().Root.Find("[ng-view]")
	sel2 := sel.PrevAll().End()
	AssertEqual(t, sel, sel2)
}

func TestPrevAll2(t *testing.T) {
	sel := Doc().Root.Find(".pvk-gutter").PrevAll()
	AssertLength(t, sel.Nodes, 6)
}

func TestPrevAllFiltered(t *testing.T) {
	sel := Doc().Root.Find(".pvk-gutter").PrevAllFiltered(".pvk-content")
	AssertLength(t, sel.Nodes, 3)
}

func TestPrevAllFilteredRollback(t *testing.T) {
	sel := Doc().Root.Find(".pvk-gutter")
	sel2 := sel.PrevAllFiltered(".pvk-content").End()
	AssertEqual(t, sel, sel2)
}
