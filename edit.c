/*
 * edit.c - Execute an external editor to edit files
 *
 * Copyright (c) 2017-2019 Rudá Moura. All rights reserved.
 * 
 * Permission to use, copy, modify, distribute, and sell this software and its
 * documentation for any purpose is hereby granted without fee, provided that
 * the above copyright notice appear in all copies and that both that
 * copyright notice and this permission notice appear in supporting
 * documentation.  No representations are made about the suitability of this
 * software for any purpose.  It is provided "as is" without express or
 * implied warranty.
 */

#include <errno.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

int
main(int argc, char *argv[])
{
	char *progname = argv[0];
	char *editor = getenv("EDITOR");
	if (editor == NULL)
		editor = "vi";
	argv[0] = editor;
	int res = execvp(editor, argv);
	if (res == -1) {
		printf("%s: exec error: %s '%s'\n", progname, strerror(errno), editor);
		return 127;
	}
	return 0;  /* never reached */
}
