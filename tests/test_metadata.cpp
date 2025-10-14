#include <stdio.h>
#include <libraw/libraw.h>

int main(int argc, char *argv[]) {
    if (argc != 2) {
        printf("Usage: %s <raw_file>\n", argv[0]);
        return 1;
    }

    const char *filename = argv[1];
    LibRaw processor;
    int ret;

    if ((ret = processor.open_file(filename)) != LIBRAW_SUCCESS) {
        printf("Error opening file: %s\n", libraw_strerror(ret));
        return 1;
    }

    if ((ret = processor.unpack()) != LIBRAW_SUCCESS) {
        printf("Error unpacking file: %s\n", libraw_strerror(ret));
        return 1;
    }

    // Use the correct type
    libraw_iparams_t *idata = &processor.imgdata.idata;
	libraw_image_sizes_t *sizes = &processor.imgdata.sizes;

	// idata
    printf("Make: %s\n", idata->make);
    printf("Model: %s\n", idata->model);
    printf("Software: %s\n", idata->software);

    printf("MakerIndex: %d\n", idata->maker_index);
    printf("RawCount: %d\n", idata->raw_count);
    printf("IsFoveon: %d\n", idata->is_foveon);
    printf("DngVersion: %d\n", idata->dng_version);

	printf("Colors: %d\n", idata->colors);
	printf("Color descriptions: %s\n", idata->cdesc);

	// size
	printf("RawHeight: %d\n", sizes->raw_height);
	printf("RawWidth: %d\n", sizes->raw_width);
	printf("Height: %d\n", sizes->height);
	printf("Width: %d\n", sizes->width);
	printf("IHeight: %d\n", sizes->iheight);
	printf("IWidth: %d\n", sizes->iwidth);

    processor.recycle();
    return 0;
}
