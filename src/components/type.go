package components

import (
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/textinput"
)

// Type representing the mode of the panel
type panelMode uint

// Type representing the focus type of the file panel
type filePanelFocusType uint

// Type representing the state of a process
type processState int

// Type representing the type of focused panel
type focusPanelType int

// Type representing the type of item
type itemType int

type warnType int

const (
	confirmDeleteItem warnType = iota
)

// Constants for new file or new directory
const (
	newFile itemType = iota
	newDirectory
)

// Constants for panel with no focus
const (
	nonePanelFocus focusPanelType = iota
	processBarFocus
	sidebarFocus
	metadataFocus
)

// Constants for file panel with no focus
const (
	noneFocus filePanelFocusType = iota
	secondFocus
	focus
)

// Constants for select mode or browser mode
const (
	selectMode panelMode = iota
	browserMode
)

// Constants for operation, success, cancel, failure
const (
	inOperation processState = iota
	successful
	cancel
	failure
)

// Main model
type model struct {
	fileModel           fileModel
	sidebarModel        sidebarModel
	processBarModel     processBarModel
	focusPanel          focusPanelType
	copyItems           copyItems
	typingModal         typingModal
	warnModal           warnModal
	fileMetaData        fileMetadata
	firstTextInput      bool
	toggleDotFile       bool
	editorMode          bool
	filePanelFocusIndex int
	mainPanelHeight     int
	fullWidth           int
	fullHeight          int
}

// Modal

type warnModal struct {
	open     bool
	warnType warnType
	title    string
	content  string
}

type typingModal struct {
	location  string
	open      bool
	itemType  itemType
	textInput textinput.Model
}

// File metadata
type fileMetadata struct {
	metaData    [][2]string
	renderIndex int
}

// Copied items
type copyItems struct {
	items         []string
	originalPanel originalPanel
	cut           bool
}

// Original panel
type originalPanel struct {
	index    int
	location string
}

/* FILE WINDOWS TYPE START*/
// Model for file windows
type fileModel struct {
	filePanels   []filePanel
	width        int
	renaming     bool
	maxFilePanel int
}

// Panel representing a file
type filePanel struct {
	cursor             int
	render             int
	focusType          filePanelFocusType
	location           string
	panelMode          panelMode
	selected           []string
	element            []element
	directoryRecord    map[string]directoryRecord
	rename             textinput.Model
	renaming           bool
	searchBar          textinput.Model
	lastTimeGetElement time.Time
}

// Record for directory navigation
type directoryRecord struct {
	directoryCursor int
	directoryRender int
}

// Element within a file panel
type element struct {
	name      string
	location  string
	directory bool
	matchRate float64
	metaData  [][2]string
}

/* FILE WINDOWS TYPE END*/

/* SIDE BAR COMPONENTS TYPE START*/
// Model for sidebar components
type sidebarModel struct {
	directories []directory
	// wellKnownModel []directory
	// pinnedModel    []directory
	// disksModel     []directory
	cursor int
}

type directory struct {
	location string
	name     string
}

/* SIDE BAR COMPONENTS TYPE END*/

/*PROCESS BAR COMPONENTS TYPE START*/

// Model for process bar components
type processBarModel struct {
	render      int
	cursor      int
	processList []string
	process     map[string]process
}

// Model for an individual process
type process struct {
	name     string
	progress progress.Model
	state    processState
	total    int
	done     int
	doneTime time.Time
}

// Message for process bar
type channelMessage struct {
	messageId       string
	processNewState process
	returnWarnModal bool
	warnModal       warnModal
	loadMetadata    bool
	metadata        [][2]string
}

/*PROCESS BAR COMPONENTS TYPE END*/

// Style for icons
type iconStyle struct {
	icon  string
	color string
}

type editorFinishedMsg struct{ err error }
