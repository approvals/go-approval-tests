package reporters

import (
	"os"
	"runtime"
	"strings"
)

func expandProgramFiles(path string) string {
	return strings.ReplaceAll(strings.Replace(path, "{ProgramFiles}", os.Getenv("ProgramFiles"), 1), "\\", "/")
}


type diffMergeMac struct{}

func NewDiffMergeMacReporter() Reporter {
	return &diffMergeMac{}
}

func (s *diffMergeMac) Report(approved, received string) bool {
	if runtime.GOOS != goosDarwin {
		return false
	}
	programName := "/Applications/DiffMerge.app/Contents/MacOS/DiffMerge"
	args := []string{"--nosplash", received, approved}
	return launchProgram(programName, approved, args...)
}

type beyondCompareMac struct{}

func NewBeyondCompareMacReporter() Reporter {
	return &beyondCompareMac{}
}

func (s *beyondCompareMac) Report(approved, received string) bool {
	if runtime.GOOS != goosDarwin {
		return false
	}
	programName := "/Applications/Beyond Compare.app/Contents/MacOS/bcomp"
	args := []string{received, approved}
	return launchProgram(programName, approved, args...)
}

type kaleidoscopeMac struct{}

func NewKaleidoscopeMacReporter() Reporter {
	return &kaleidoscopeMac{}
}

func (s *kaleidoscopeMac) Report(approved, received string) bool {
	if runtime.GOOS != goosDarwin {
		return false
	}
	programName := "/Applications/Kaleidoscope.app/Contents/MacOS/ksdiff"
	args := []string{received, approved}
	return launchProgram(programName, approved, args...)
}

type kaleidoscope3Mac struct{}

func NewKaleidoscope3MacReporter() Reporter {
	return &kaleidoscope3Mac{}
}

func (s *kaleidoscope3Mac) Report(approved, received string) bool {
	if runtime.GOOS != goosDarwin {
		return false
	}
	programName := "/usr/local/bin/ksdiff"
	args := []string{received, approved}
	return launchProgram(programName, approved, args...)
}

type kdiff3Mac struct{}

func NewKdiff3MacReporter() Reporter {
	return &kdiff3Mac{}
}

func (s *kdiff3Mac) Report(approved, received string) bool {
	if runtime.GOOS != goosDarwin {
		return false
	}
	programName := "/Applications/kdiff3.app/Contents/MacOS/kdiff3"
	args := []string{received, approved, "-m"}
	return launchProgram(programName, approved, args...)
}

type p4mergeMac struct{}

func NewP4mergeMacReporter() Reporter {
	return &p4mergeMac{}
}

func (s *p4mergeMac) Report(approved, received string) bool {
	if runtime.GOOS != goosDarwin {
		return false
	}
	programName := "/Applications/p4merge.app/Contents/MacOS/p4merge"
	args := []string{received, approved}
	return launchProgram(programName, approved, args...)
}

type tkDiffMac struct{}

func NewTkDiffMacReporter() Reporter {
	return &tkDiffMac{}
}

func (s *tkDiffMac) Report(approved, received string) bool {
	if runtime.GOOS != goosDarwin {
		return false
	}
	programName := "/Applications/TkDiff.app/Contents/MacOS/tkdiff"
	args := []string{received, approved}
	return launchProgram(programName, approved, args...)
}

type visualStudioCodeMac struct{}

func NewVisualStudioCodeMacReporter() Reporter {
	return &visualStudioCodeMac{}
}

func (s *visualStudioCodeMac) Report(approved, received string) bool {
	if runtime.GOOS != goosDarwin {
		return false
	}
	programName := "/Applications/Visual Studio Code.app/Contents/Resources/app/bin/code"
	args := []string{"-d", received, approved}
	return launchProgram(programName, approved, args...)
}

type araxisMergeMac struct{}

func NewAraxisMergeMacReporter() Reporter {
	return &araxisMergeMac{}
}

func (s *araxisMergeMac) Report(approved, received string) bool {
	if runtime.GOOS != goosDarwin {
		return false
	}
	programName := "/Applications/Araxis Merge.app/Contents/Utilities/compare"
	args := []string{received, approved}
	return launchProgram(programName, approved, args...)
}

type beyondCompare3Windows struct{}

func NewBeyondCompare3WindowsReporter() Reporter {
	return &beyondCompare3Windows{}
}

func (s *beyondCompare3Windows) Report(approved, received string) bool {
	if runtime.GOOS != goosWindows {
		return false
	}
	programName := expandProgramFiles("{ProgramFiles}Beyond Compare 3\\BCompare.exe")
	args := []string{received, approved}
	return launchProgram(programName, approved, args...)
}

type beyondCompare4Windows struct{}

func NewBeyondCompare4WindowsReporter() Reporter {
	return &beyondCompare4Windows{}
}

func (s *beyondCompare4Windows) Report(approved, received string) bool {
	if runtime.GOOS != goosWindows {
		return false
	}
	programName := expandProgramFiles("{ProgramFiles}Beyond Compare 4\\BCompare.exe")
	args := []string{received, approved}
	return launchProgram(programName, approved, args...)
}

type beyondCompare5Windows struct{}

func NewBeyondCompare5WindowsReporter() Reporter {
	return &beyondCompare5Windows{}
}

func (s *beyondCompare5Windows) Report(approved, received string) bool {
	if runtime.GOOS != goosWindows {
		return false
	}
	programName := expandProgramFiles("{ProgramFiles}Beyond Compare 5\\BCompare.exe")
	args := []string{received, approved}
	return launchProgram(programName, approved, args...)
}

type tortoiseImageDiffWindows struct{}

func NewTortoiseImageDiffWindowsReporter() Reporter {
	return &tortoiseImageDiffWindows{}
}

func (s *tortoiseImageDiffWindows) Report(approved, received string) bool {
	if runtime.GOOS != goosWindows {
		return false
	}
	programName := expandProgramFiles("{ProgramFiles}TortoiseSVN\\bin\\TortoiseIDiff.exe")
	args := []string{"/left:" + received, "/right:" + approved}
	return launchProgram(programName, approved, args...)
}

type tortoiseTextDiffWindows struct{}

func NewTortoiseTextDiffWindowsReporter() Reporter {
	return &tortoiseTextDiffWindows{}
}

func (s *tortoiseTextDiffWindows) Report(approved, received string) bool {
	if runtime.GOOS != goosWindows {
		return false
	}
	programName := expandProgramFiles("{ProgramFiles}TortoiseSVN\\bin\\TortoiseMerge.exe")
	args := []string{received, approved}
	return launchProgram(programName, approved, args...)
}

type tortoiseGitImageDiffWindows struct{}

func NewTortoiseGitImageDiffWindowsReporter() Reporter {
	return &tortoiseGitImageDiffWindows{}
}

func (s *tortoiseGitImageDiffWindows) Report(approved, received string) bool {
	if runtime.GOOS != goosWindows {
		return false
	}
	programName := expandProgramFiles("{ProgramFiles}TortoiseGIT\\bin\\TortoiseGitIDiff.exe")
	args := []string{"/left:" + received, "/right:" + approved}
	return launchProgram(programName, approved, args...)
}

type tortoiseGitTextDiffWindows struct{}

func NewTortoiseGitTextDiffWindowsReporter() Reporter {
	return &tortoiseGitTextDiffWindows{}
}

func (s *tortoiseGitTextDiffWindows) Report(approved, received string) bool {
	if runtime.GOOS != goosWindows {
		return false
	}
	programName := expandProgramFiles("{ProgramFiles}TortoiseGIT\\bin\\TortoiseGitMerge.exe")
	args := []string{received, approved}
	return launchProgram(programName, approved, args...)
}

type winMergeReporterWindows struct{}

func NewWinMergeReporterWindowsReporter() Reporter {
	return &winMergeReporterWindows{}
}

func (s *winMergeReporterWindows) Report(approved, received string) bool {
	if runtime.GOOS != goosWindows {
		return false
	}
	programName := expandProgramFiles("{ProgramFiles}WinMerge\\WinMergeU.exe")
	args := []string{received, approved}
	return launchProgram(programName, approved, args...)
}

type araxisMergeWindows struct{}

func NewAraxisMergeWindowsReporter() Reporter {
	return &araxisMergeWindows{}
}

func (s *araxisMergeWindows) Report(approved, received string) bool {
	if runtime.GOOS != goosWindows {
		return false
	}
	programName := expandProgramFiles("{ProgramFiles}Araxis\\Araxis Merge\\Compare.exe")
	args := []string{received, approved}
	return launchProgram(programName, approved, args...)
}

type codeCompareWindows struct{}

func NewCodeCompareWindowsReporter() Reporter {
	return &codeCompareWindows{}
}

func (s *codeCompareWindows) Report(approved, received string) bool {
	if runtime.GOOS != goosWindows {
		return false
	}
	programName := expandProgramFiles("{ProgramFiles}Devart\\Code Compare\\CodeCompare.exe")
	args := []string{received, approved}
	return launchProgram(programName, approved, args...)
}

type kdiff3Windows struct{}

func NewKdiff3WindowsReporter() Reporter {
	return &kdiff3Windows{}
}

func (s *kdiff3Windows) Report(approved, received string) bool {
	if runtime.GOOS != goosWindows {
		return false
	}
	programName := expandProgramFiles("{ProgramFiles}KDiff3\\kdiff3.exe")
	args := []string{received, approved}
	return launchProgram(programName, approved, args...)
}

type visualStudioCodeWindows struct{}

func NewVisualStudioCodeWindowsReporter() Reporter {
	return &visualStudioCodeWindows{}
}

func (s *visualStudioCodeWindows) Report(approved, received string) bool {
	if runtime.GOOS != goosWindows {
		return false
	}
	programName := expandProgramFiles("{ProgramFiles}Microsoft VS Code\\Code.exe")
	args := []string{"-d", received, approved}
	return launchProgram(programName, approved, args...)
}

type diffMergeLinux struct{}

func NewDiffMergeLinuxReporter() Reporter {
	return &diffMergeLinux{}
}

func (s *diffMergeLinux) Report(approved, received string) bool {
	if runtime.GOOS != goosLinux {
		return false
	}
	programName := "/usr/bin/diffmerge"
	args := []string{"--nosplash", received, approved}
	return launchProgram(programName, approved, args...)
}

type meldMergeLinux struct{}

func NewMeldMergeLinuxReporter() Reporter {
	return &meldMergeLinux{}
}

func (s *meldMergeLinux) Report(approved, received string) bool {
	if runtime.GOOS != goosLinux {
		return false
	}
	programName := "/usr/bin/meld"
	args := []string{received, approved}
	return launchProgram(programName, approved, args...)
}

type kdiff3Linux struct{}

func NewKdiff3LinuxReporter() Reporter {
	return &kdiff3Linux{}
}

func (s *kdiff3Linux) Report(approved, received string) bool {
	if runtime.GOOS != goosLinux {
		return false
	}
	programName := "/usr/bin/kdiff3"
	args := []string{received, approved, "-m"}
	return launchProgram(programName, approved, args...)
}

type diffCommandLineLinux struct{}

func NewDiffCommandLineLinuxReporter() Reporter {
	return &diffCommandLineLinux{}
}

func (s *diffCommandLineLinux) Report(approved, received string) bool {
	if runtime.GOOS != goosLinux {
		return false
	}
	programName := "/usr/bin/diff"
	args := []string{"-u", received, approved}
	return launchProgram(programName, approved, args...)
}

type diffCommandLineMac struct{}

func NewDiffCommandLineMacReporter() Reporter {
	return &diffCommandLineMac{}
}

func (s *diffCommandLineMac) Report(approved, received string) bool {
	if runtime.GOOS != goosDarwin {
		return false
	}
	programName := "/usr/bin/diff"
	args := []string{"-u", received, approved}
	return launchProgram(programName, approved, args...)
}

type sublimeMergeMac struct{}

func NewSublimeMergeMacReporter() Reporter {
	return &sublimeMergeMac{}
}

func (s *sublimeMergeMac) Report(approved, received string) bool {
	if runtime.GOOS != goosDarwin {
		return false
	}
	programName := "/Applications/Sublime Merge.app/Contents/SharedSupport/bin/smerge"
	args := []string{"mergetool", received, approved}
	return launchProgram(programName, approved, args...)
}

type sublimeMergeWindows struct{}

func NewSublimeMergeWindowsReporter() Reporter {
	return &sublimeMergeWindows{}
}

func (s *sublimeMergeWindows) Report(approved, received string) bool {
	if runtime.GOOS != goosWindows {
		return false
	}
	programName := expandProgramFiles("{ProgramFiles}Sublime Merge\\smerge.exe")
	args := []string{"mergetool", received, approved}
	return launchProgram(programName, approved, args...)
}

type sublimeMergeLinux struct{}

func NewSublimeMergeLinuxReporter() Reporter {
	return &sublimeMergeLinux{}
}

func (s *sublimeMergeLinux) Report(approved, received string) bool {
	if runtime.GOOS != goosLinux {
		return false
	}
	programName := "/usr/bin/smerge"
	args := []string{"mergetool", received, approved}
	return launchProgram(programName, approved, args...)
}

type diffToolOnMac struct{}

func NewDiffToolOnMacReporter() Reporter {
	return NewFirstWorkingReporter(
		NewDiffMergeMacReporter(),
		NewBeyondCompareMacReporter(),
		NewKaleidoscopeMacReporter(),
		NewKaleidoscope3MacReporter(),
		NewKdiff3MacReporter(),
		NewP4mergeMacReporter(),
		NewTkDiffMacReporter(),
		NewVisualStudioCodeMacReporter(),
		NewAraxisMergeMacReporter(),
		NewDiffCommandLineMacReporter(),
		NewSublimeMergeMacReporter(),
	)
}

func (s *diffToolOnMac) Report(approved, received string) bool {
	if runtime.GOOS != goosDarwin {
		return false
	}
	return NewFirstWorkingReporter(
		NewDiffMergeMacReporter(),
		NewBeyondCompareMacReporter(),
		NewKaleidoscopeMacReporter(),
		NewKaleidoscope3MacReporter(),
		NewKdiff3MacReporter(),
		NewP4mergeMacReporter(),
		NewTkDiffMacReporter(),
		NewVisualStudioCodeMacReporter(),
		NewAraxisMergeMacReporter(),
		NewDiffCommandLineMacReporter(),
		NewSublimeMergeMacReporter(),
	).Report(approved, received)
}

type diffToolOnWindows struct{}

func NewDiffToolOnWindowsReporter() Reporter {
	return NewFirstWorkingReporter(
		NewBeyondCompare3WindowsReporter(),
		NewBeyondCompare4WindowsReporter(),
		NewBeyondCompare5WindowsReporter(),
		NewTortoiseImageDiffWindowsReporter(),
		NewTortoiseTextDiffWindowsReporter(),
		NewTortoiseGitImageDiffWindowsReporter(),
		NewTortoiseGitTextDiffWindowsReporter(),
		NewWinMergeReporterWindowsReporter(),
		NewAraxisMergeWindowsReporter(),
		NewCodeCompareWindowsReporter(),
		NewKdiff3WindowsReporter(),
		NewVisualStudioCodeWindowsReporter(),
		NewSublimeMergeWindowsReporter(),
	)
}

func (s *diffToolOnWindows) Report(approved, received string) bool {
	if runtime.GOOS != goosWindows {
		return false
	}
	return NewFirstWorkingReporter(
		NewBeyondCompare3WindowsReporter(),
		NewBeyondCompare4WindowsReporter(),
		NewBeyondCompare5WindowsReporter(),
		NewTortoiseImageDiffWindowsReporter(),
		NewTortoiseTextDiffWindowsReporter(),
		NewTortoiseGitImageDiffWindowsReporter(),
		NewTortoiseGitTextDiffWindowsReporter(),
		NewWinMergeReporterWindowsReporter(),
		NewAraxisMergeWindowsReporter(),
		NewCodeCompareWindowsReporter(),
		NewKdiff3WindowsReporter(),
		NewVisualStudioCodeWindowsReporter(),
		NewSublimeMergeWindowsReporter(),
	).Report(approved, received)
}

type diffToolOnLinux struct{}

func NewDiffToolOnLinuxReporter() Reporter {
	return NewFirstWorkingReporter(
		NewDiffMergeLinuxReporter(),
		NewMeldMergeLinuxReporter(),
		NewKdiff3LinuxReporter(),
		NewDiffCommandLineLinuxReporter(),
		NewSublimeMergeLinuxReporter(),
	)
}

func (s *diffToolOnLinux) Report(approved, received string) bool {
	if runtime.GOOS != goosLinux {
		return false
	}
	return NewFirstWorkingReporter(
		NewDiffMergeLinuxReporter(),
		NewMeldMergeLinuxReporter(),
		NewKdiff3LinuxReporter(),
		NewDiffCommandLineLinuxReporter(),
		NewSublimeMergeLinuxReporter(),
	).Report(approved, received)
}

type araxisMergeGroup struct{}

func NewAraxisMergeGroupReporter() Reporter {
	return NewFirstWorkingReporter(
		NewAraxisMergeMacReporter(),
		NewAraxisMergeWindowsReporter(),
	)
}

func (s *araxisMergeGroup) Report(approved, received string) bool {
	return NewFirstWorkingReporter(
		NewAraxisMergeMacReporter(),
		NewAraxisMergeWindowsReporter(),
	).Report(approved, received)
}

type beyondCompareGroup struct{}

func NewBeyondCompareGroupReporter() Reporter {
	return NewFirstWorkingReporter(
		NewBeyondCompareMacReporter(),
		NewBeyondCompare3WindowsReporter(),
		NewBeyondCompare4WindowsReporter(),
		NewBeyondCompare5WindowsReporter(),
	)
}

func (s *beyondCompareGroup) Report(approved, received string) bool {
	return NewFirstWorkingReporter(
		NewBeyondCompareMacReporter(),
		NewBeyondCompare3WindowsReporter(),
		NewBeyondCompare4WindowsReporter(),
		NewBeyondCompare5WindowsReporter(),
	).Report(approved, received)
}

type diffCommandLineGroup struct{}

func NewDiffCommandLineGroupReporter() Reporter {
	return NewFirstWorkingReporter(
		NewDiffCommandLineLinuxReporter(),
		NewDiffCommandLineMacReporter(),
	)
}

func (s *diffCommandLineGroup) Report(approved, received string) bool {
	return NewFirstWorkingReporter(
		NewDiffCommandLineLinuxReporter(),
		NewDiffCommandLineMacReporter(),
	).Report(approved, received)
}

type diffMergeGroup struct{}

func NewDiffMergeGroupReporter() Reporter {
	return NewFirstWorkingReporter(
		NewDiffMergeMacReporter(),
		NewDiffMergeLinuxReporter(),
	)
}

func (s *diffMergeGroup) Report(approved, received string) bool {
	return NewFirstWorkingReporter(
		NewDiffMergeMacReporter(),
		NewDiffMergeLinuxReporter(),
	).Report(approved, received)
}

type kaleidoscopeGroup struct{}

func NewKaleidoscopeGroupReporter() Reporter {
	return NewFirstWorkingReporter(
		NewKaleidoscopeMacReporter(),
		NewKaleidoscope3MacReporter(),
	)
}

func (s *kaleidoscopeGroup) Report(approved, received string) bool {
	return NewFirstWorkingReporter(
		NewKaleidoscopeMacReporter(),
		NewKaleidoscope3MacReporter(),
	).Report(approved, received)
}

type kdiff3Group struct{}

func NewKdiff3GroupReporter() Reporter {
	return NewFirstWorkingReporter(
		NewKdiff3MacReporter(),
		NewKdiff3WindowsReporter(),
		NewKdiff3LinuxReporter(),
	)
}

func (s *kdiff3Group) Report(approved, received string) bool {
	return NewFirstWorkingReporter(
		NewKdiff3MacReporter(),
		NewKdiff3WindowsReporter(),
		NewKdiff3LinuxReporter(),
	).Report(approved, received)
}

type sublimeMergeGroup struct{}

func NewSublimeMergeGroupReporter() Reporter {
	return NewFirstWorkingReporter(
		NewSublimeMergeMacReporter(),
		NewSublimeMergeWindowsReporter(),
		NewSublimeMergeLinuxReporter(),
	)
}

func (s *sublimeMergeGroup) Report(approved, received string) bool {
	return NewFirstWorkingReporter(
		NewSublimeMergeMacReporter(),
		NewSublimeMergeWindowsReporter(),
		NewSublimeMergeLinuxReporter(),
	).Report(approved, received)
}

type tortoiseGroup struct{}

func NewTortoiseGroupReporter() Reporter {
	return NewFirstWorkingReporter(
		NewTortoiseImageDiffWindowsReporter(),
		NewTortoiseTextDiffWindowsReporter(),
	)
}

func (s *tortoiseGroup) Report(approved, received string) bool {
	return NewFirstWorkingReporter(
		NewTortoiseImageDiffWindowsReporter(),
		NewTortoiseTextDiffWindowsReporter(),
	).Report(approved, received)
}

type tortoiseGitGroup struct{}

func NewTortoiseGitGroupReporter() Reporter {
	return NewFirstWorkingReporter(
		NewTortoiseGitImageDiffWindowsReporter(),
		NewTortoiseGitTextDiffWindowsReporter(),
	)
}

func (s *tortoiseGitGroup) Report(approved, received string) bool {
	return NewFirstWorkingReporter(
		NewTortoiseGitImageDiffWindowsReporter(),
		NewTortoiseGitTextDiffWindowsReporter(),
	).Report(approved, received)
}

type visualStudioCodeGroup struct{}

func NewVisualStudioCodeGroupReporter() Reporter {
	return NewFirstWorkingReporter(
		NewVisualStudioCodeMacReporter(),
		NewVisualStudioCodeWindowsReporter(),
	)
}

func (s *visualStudioCodeGroup) Report(approved, received string) bool {
	return NewFirstWorkingReporter(
		NewVisualStudioCodeMacReporter(),
		NewVisualStudioCodeWindowsReporter(),
	).Report(approved, received)
}
