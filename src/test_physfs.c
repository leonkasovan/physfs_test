#include <stdio.h>
#include <stdlib.h>
#include <physfs.h>

int main(int argc, char **argv)
{
    // Initialize the PhysicsFS library
    if (PHYSFS_init(argv[0]) == 0)
    {
        fprintf(stderr, "Failed to initialize PhysicsFS: %s\n", PHYSFS_getErrorByCode(PHYSFS_getLastErrorCode()));
        return 1;
    }

    const PHYSFS_ArchiveInfo **i;
    for (i = PHYSFS_supportedArchiveTypes(); *i != NULL; i++)
    {
        printf("Supported archive: [%s], which is [%s].\n", (*i)->extension, (*i)->description);
    }

    // Mount an archive (e.g., a ZIP file)
    if (PHYSFS_mount("example.zip", "/", 1) == 0)
    {
        fprintf(stderr, "Failed to mount archive: %s\n", PHYSFS_getErrorByCode(PHYSFS_getLastErrorCode()));
        PHYSFS_deinit();
        return 1;
    }

    // Mount regular directory
    if (PHYSFS_mount(".", "/", 1) == 0)
    {
        fprintf(stderr, "Failed to mount directory: %s\n", PHYSFS_getErrorByCode(PHYSFS_getLastErrorCode()));
        PHYSFS_deinit();
        return 1;
    }

    // Open a file for reading
    PHYSFS_File *file = PHYSFS_openRead("packages/physfs/physfs.go");
    if (file == NULL)
    {
        fprintf(stderr, "Failed to open file: %s\n", PHYSFS_getErrorByCode(PHYSFS_getLastErrorCode()));
        PHYSFS_deinit();
        return 1;
    }

    // Read the file contents
    PHYSFS_sint64 fileSize = PHYSFS_fileLength(file);
    char *buffer = (char *)malloc(fileSize + 1);
    if (buffer == NULL)
    {
        fprintf(stderr, "Failed to allocate memory\n");
        PHYSFS_close(file);
        PHYSFS_deinit();
        return 1;
    }
    PHYSFS_readBytes(file, buffer, fileSize);
    buffer[fileSize] = '\0'; // Null-terminate the string

    // Print the file contents
    printf("File contents: %s\n", buffer);

    // Unmount the archive
    if (PHYSFS_unmount("example.zip") == 0)
    {
        fprintf(stderr, "Failed to unmount archive: %s\n", PHYSFS_getErrorByCode(PHYSFS_getLastErrorCode()));
        free(buffer);
        PHYSFS_close(file);
        PHYSFS_deinit();
        return 1;
    }

    // Unmount the directory
    if (PHYSFS_unmount(".") == 0)
    {
        fprintf(stderr, "Failed to unmount directory: %s\n", PHYSFS_getErrorByCode(PHYSFS_getLastErrorCode()));
        free(buffer);
        PHYSFS_close(file);
        PHYSFS_deinit();
        return 1;
    }

    // Clean up
    free(buffer);
    PHYSFS_close(file);
    PHYSFS_deinit();

    return 0;
}