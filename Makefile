CC=		cc
CFLAGS=		-Wall -pedantic -Os -pipe

PREFIX=		/usr/local
BINDIR=		$(PREFIX)/bin
MANDIR=		$(PREFIX)/share/man

all: edit

install: edit
	test -d $(DESTDIR)$(BINDIR) || install -d $(DESTDIR)$(BINDIR) || true
	install -m 755 edit $(DESTDIR)$(BINDIR)
	test -d $(DESTDIR)$(MANDIR)/man1 || install -d $(DESTDIR)$(MANDIR)/man1 || true
	install -m 644 edit.1 $(DESTDIR)$(MANDIR)/man1

uninstall:
	rm -f $(BINDIR)/edit
	rm -f $(MANDIR)/man1/edit.1

edit.ps: edit.1
	groff -man -Tps edit.1 > $@

clean:
	rm -f *~ edit edit.ps

PHONY: clean
