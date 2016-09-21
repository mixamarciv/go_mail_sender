CALL "%~dp0/set_path.bat"



::@CLS
::@pause

@echo === install ===================================================================
go get -u "github.com/mixamarciv/gofncstd3000"
go get -u "github.com/go-gomail/gomail"
go get -u "github.com/jessevdk/go-flags"

go install

@echo ==== end ======================================================================
@PAUSE
