package controls

type NavType int

const (
	Heading NavType = iota
	NavHrefBlank
	NavButton
	NavHref
	NavHRule
)

type NavMenuItem struct {
	Type   NavType
	Title  string
	HREF   string
	Target string
}

func NewNavHRule() NavMenuItem {
	nmi := NavMenuItem{
		Type: NavHRule,
	}
	return nmi
}

func NewNavHref(title string, href string) NavMenuItem {
	nmi := NavMenuItem{
		Type:  NavHref,
		Title: title,
		HREF:  href,
	}
	return nmi
}

func NewNavButton(title string, href string, target string) NavMenuItem {
	nmi := NavMenuItem{
		Type:   NavButton,
		Title:  title,
		HREF:   href,
		Target: target,
	}
	return nmi
}

func NewNavHrefBlank(title string, href string) NavMenuItem {
	nmi := NavMenuItem{
		Type:  NavHrefBlank,
		Title: title,
		HREF:  href,
	}
	return nmi
}

func NewHeading(title string) NavMenuItem {
	nmi := NavMenuItem{
		Type:  Heading,
		Title: title,
	}

	return nmi
}
