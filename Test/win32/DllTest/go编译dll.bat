@echo off
set filepath=%1
set dllpath=%~dp1%~n1.dll

IF NOT EXIST "%filepath%" (
	goto err
) ELSE (
	goto make
)

:err
echo �ļ������ڣ�
set/p exepath=����ק�ļ�����������ļ�·��Ȼ���»س�
IF NOT EXIST "%filepath%" (
	goto err
) ELSE (
	goto make
)

:make
echo ���� %~nx1 -^> %~n1.dll
call go build -ldflags "-s -w" -buildmode=c-shared -o "%dllpath%" "%filepath%"
echo �������!���»س��رմ���

:end
pause