<template>
    <div>
        <p>
            This is a simple url-shorter on your localhost.<br />
            Made with Vue.js, Go and Redis.
        </p>
    </div>
    <div>
        <input placeholder="Paste here your url..." type="text" v-model="URL"/>
        <button @click="short">Short</button>
    </div>
</template>

<script>
import { ref, defineComponent } from 'vue';

export default defineComponent({
    name: 'SaveBar',
    emits: ['getUrl'],
    setup(props, {emit}){
        let URL = ref("")
        const short = () => {
            fetch('http://localhost:8888/save', {
                method: 'POST',
                headers: {'Content-Type': 'application/json'},
                body: JSON.stringify({ url: URL.value })
            })
            .then(response => { response.json().then(json => emit('getUrl', json.url))})
            .catch(err => console.log(err))
        }
        return {URL, short}
    }
})
</script>

<style scoped>
div {
    display: flex;
    flex-direction: row;
    align-items: center;
}

input {
    height: 30px;
    width: 500px;
    margin: 0 5px;
    border-color: #0C0F0A;
    border-radius: 2px;
}

button {
    height: 35px;
    width: 100px;
    margin: 0 5px;
    border-color: #0C0F0A;
    border-radius: 5px;
    transition: all 0.5s ease;
}

button:hover {
    transform: scale(105%);
    background-color: #FBFF12;
    font: bold;
    color: #0C0F0A; 
}
</style>