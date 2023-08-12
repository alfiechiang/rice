<template>
    <div class="tag">
        <div id="tagArea">
            <el-tag v-for="tag in dynamicTags" :key="tag" class="mx-1" closable :disable-transitions="false"
                @close="handleClose(tag.name)" :type="tag.type" @click="handleClick(tag)" style="margin-right:2px">
                <a>{{ tag.name }}</a>
            </el-tag>
        </div>
    </div>
</template>
<script lang="ts" setup >
import { nextTick, ref, watch } from 'vue'
import { useRoute, useRouter } from "vue-router";
import { type TagItem } from './types';
const router = useRouter();
const dynamicTags = ref<TagItem[]>([])
const currentRoute = useRoute();

watch(
    () => currentRoute.path,
    (route) => {
        addTag()
    }
);

const addTag = () => {
    let contains = false;
    dynamicTags.value.forEach((item, v) => {
        if (item.name == currentRoute.meta.title) {
            contains = true;
        }
    })
    if (!contains) {
        dynamicTags.value.forEach((item, v) => {
            item.type = 'info';
        })
        dynamicTags.value.push({ name: currentRoute.meta.title as string, path: currentRoute.path, type: '' })
    }
}

const handleClose = (tagName: string) => {
    const indexToRemove = dynamicTags.value.findIndex(item => item.name === tagName);
    if (indexToRemove !== -1) {
        dynamicTags.value.splice(indexToRemove, 1);
    }
}

const handleClick = (item: TagItem) => {
    dynamicTags.value.forEach((item) => {item.type = 'info';})
    item.type = '';
    router.push({
        path: item.path,
    });
}

</script>

<style scoped lang="scss">
.tag {
    width: 100%;
    height: 30px;
    border-bottom: 1px solid #E5E7ED;
    display: flex;
    align-items: center;

    #tagArea {
        a:hover {
            cursor: pointer;
        }
    }
}
</style>
  