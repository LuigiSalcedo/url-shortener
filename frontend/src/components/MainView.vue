<template>
    <TitleAuthor/>
    <SaveBar @getUrl="updateUrl"/>
    <ResultUrl :url="url"/>
</template>

<script setup>
import TitleAuthor from './TitleAuthor.vue'
import SaveBar from './SaveBar.vue';
import ResultUrl from './ResultUrl.vue';
import { ref } from 'vue';


let url = ref("")

const updateUrl = (newUrl) => {
    url.value = newUrl
}
</script>

<style>
</style>