<template>
    <div class="common-layout">
        <el-container style="height:100%">
            <el-aside width="200px">
                <leftmenu></leftmenu>
            </el-aside>
            <el-container>
                <el-header height="50px">
                    <breadcrumb></breadcrumb>
                    <div class="right-header">
                        
                        <el-dropdown>

                            <div class="icon-area">
                                <img class="avatar-img" src="https://oss.youlai.tech/youlai-boot/2023/05/16/811270ef31f548af9cffc026dfc3777b.gif?imageView2/1/w/80/h/80">
                                <el-icon><ArrowDown /></el-icon>
                            </div>

                           
                            <template #dropdown>
                                <el-dropdown-menu>
                                    <el-dropdown-item>Action 1</el-dropdown-item>
                                    <el-dropdown-item>Action 2</el-dropdown-item>
                                    <el-dropdown-item>Action 3</el-dropdown-item>
                                    <el-dropdown-item disabled>Action 4</el-dropdown-item>
                                    <el-dropdown-item divided>Action 5</el-dropdown-item>
                                </el-dropdown-menu>
                            </template>
                        </el-dropdown>
                    </div>
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
        display: flex;

    }

    .el-main {
        background: #F2F3F5;
    }

    .avatar-img {
        max-width: 40px;
        border-radius: 5px;
    }

    .right-header {
        display: flex;
        justify-content: right;
        align-items: center;
        width: 10%;
        height: 100%;
    }

    .example-showcase .el-dropdown-link {
        cursor: pointer;
        color: var(--el-color-primary);
        display: flex;
        align-items: center;
    }

    .icon-area{
        width: 100px;
        height: 50px;
        display: flex;
        justify-content: right;
        align-items: center;
        img{
            margin-right: 3px;
        }
    }

}</style>