# patchwork
docker image to execute json patch

## Usage

Image: `guoyk/patchwork`

Environment variables:

   * `PATCH_SOURCE`, path to jsonpatch file
   * `PATCH_TARGET`, path to target json file
   * `PATCHWORK_HALT`, set to `true` if you don't want patchwork to exit on success, suitable for `DaemonSet`

## Credits

Guo Y.K., MIT License
