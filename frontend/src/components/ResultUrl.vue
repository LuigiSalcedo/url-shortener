<template>
    <div>
        <h4>Your new url:</h4>
        <h3> {{ props.url }}</h3>
    </div>
</template>

<script>
import { defineComponent } from 'vue'

export default defineComponent({
    name: 'ResultUrl',
    props: {
        url : {
            required: false,
            default: ''
        }
    },
    setup(props) {
        return {props}
    }
})
</script>

<style scoped>
div {
    display: flex;
}

h3, h4 {
    padding: 0 5px;
}

h3 {
    font-weight: 200;
}
</style>