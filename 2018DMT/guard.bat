@echo off

Set pn=go_build_main_go.exe

echo ������%pn%��

:LOOP
ping 127.0.0.1 -n 2 >nul
tasklist /nh|find /i "%pn%">nul
if ERRORLEVEL 1 (
    echo %date:~0,10% %time:~0,8%
    echo ���̹��ˣ�����һ��
    call "%pn%"
) else (
    Set temp=1
)
goto LOOP

:NOPROCESS


pause