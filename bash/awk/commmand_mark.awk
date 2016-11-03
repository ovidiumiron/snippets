BEGIN {
    #FS="i";
    #OFS="~";
    x=1;
    y=2;
}
#length($0) > 3 {
/i/ && /h/ {
    print $0;
    print toupper($2)
    x++;
}
END {
   print "number of matches";
   print x
}



