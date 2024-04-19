C=usr/bin/c
clear
echo "Test"
cd ~
cd ..

if test $C 
then
echo "Installed"
else
echo "clear" > $C
echo "Done!!!"
fi
