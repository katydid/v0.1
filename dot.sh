for f in *.dot; do echo "Processing $f file.."; dot -Tgif $f > ./public/$f.gif; done