<template>
    <el-tag v-for="tag in dynamicTags" :key="tag" class="mx-1" closable :disable-transitions="false"
        @close="handleClose(tag)">
        {{ tag.name }}
    </el-tag>
    <el-input v-if="inputVisible" ref="InputRef" v-model="inputValue" class="ml-1 w-20" size="small"
        @keyup.enter="handleInputConfirm" @blur="handleInputConfirm" />
</template>
  
<script lang="ts" setup>
import { nextTick, ref,watch } from 'vue'
import { ElInput } from 'element-plus'
import { useRoute } from "vue-router";


const inputValue = ref('')
const dynamicTags = ref([])
const inputVisible = ref(false)
const InputRef = ref<InstanceType<typeof ElInput>>()
const currentRoute = useRoute();


watch(
    () => currentRoute.path,
    (route) => {
        let contains=false;
        dynamicTags.value.forEach((item,v) => {
            if(item.name ==currentRoute.meta.title){
                contains=true;
            }
        })
        if(!contains){
            dynamicTags.value.push({name:currentRoute.meta.title,path:currentRoute.path})
        }

    }
);

const handleClose = (tag: string) => {
    dynamicTags.value.splice(dynamicTags.value.indexOf(tag), 1)
}

const showInput = () => {
    inputVisible.value = true
    nextTick(() => {
        InputRef.value!.input!.focus()
    })
}

const handleInputConfirm = () => {
    if (inputValue.value) {
        dynamicTags.value.push(inputValue.value)
    }
    inputVisible.value = false
    inputValue.value = ''
}
</script>
  