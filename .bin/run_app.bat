::����砥� curpath:
@FOR /f %%i IN ("%0") DO SET curpath=%~dp0
::������ �᭮��� ��६���� ���㦥���
@CALL "%curpath%/set_path.bat"


@del app.exe
@CLS

@echo === build =====================================================================
go build -o app.exe

@echo ==== start ====================================================================
app.exe --to mixamarciv@gmail.com --subject �஢�ઠ1 --body "�ਢ��3000" 

@echo ==== end ======================================================================
@PAUSE
