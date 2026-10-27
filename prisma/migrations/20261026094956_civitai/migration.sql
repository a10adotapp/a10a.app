-- CreateTable
CREATE TABLE `civitai_item_tags` (
    `id` VARCHAR(30) NOT NULL,
    `civitai_item_id` INTEGER NOT NULL,
    `civitai_tag_id` VARCHAR(191) NOT NULL,

    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE `civitai_items` (
    `id` INTEGER NOT NULL,
    `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `url` VARCHAR(191) NOT NULL,
    `base_model` VARCHAR(191) NOT NULL,
    `prompt` VARCHAR(191) NOT NULL,
    `negative_prompt` VARCHAR(191) NOT NULL,

    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE `civitai_tags` (
    `id` VARCHAR(30) NOT NULL,
    `value` VARCHAR(191) NOT NULL,

    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- AddForeignKey
ALTER TABLE `civitai_item_tags` ADD CONSTRAINT `civitai_item_tags_civitai_item_id_fkey` FOREIGN KEY (`civitai_item_id`) REFERENCES `civitai_items`(`id`) ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE `civitai_item_tags` ADD CONSTRAINT `civitai_item_tags_civitai_tag_id_fkey` FOREIGN KEY (`civitai_tag_id`) REFERENCES `civitai_tags`(`id`) ON DELETE RESTRICT ON UPDATE CASCADE;
