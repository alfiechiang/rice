<template>
    <div class="common-layout">
        <el-container style="height:100%">
            <el-aside width="200px">
                <leftmenu></leftmenu>
            </el-aside>
            <el-container>
                <el-header height="50px">
                    <breadcrumb></breadcrumb>
                </el-header>
                    <tag></tag>
                <el-main>
                    <router-view></router-view>
                </el-main>
            </el-container>
        </el-container>
    </div>
</template>
<script setup lang="ts">
import breadcrumb from '@/components/breadcrumb/index.vue'

import leftmenu from '@/components/leftMenu/index.vue'

import tag from '@/components/TagsView/index.vue'




</script>
<style  scope lang="scss">
.common-layout {
    width: 100%;
    height: 100%;

    .el-aside {
        background: #304056;
    }

    .el-header {
        border: 1px solid #E5E7ED;
    }

   

}
</style>