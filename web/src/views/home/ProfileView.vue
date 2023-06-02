<template>
    <div>
        <h2>Profile</h2>
    </div>
</template>

<script>
import { mapGetters } from "vuex";

export default {
    props: {
        userInfo: {
            type: Object,
            default: function () {
                return {};
            }
        }
    },
    computed: {
        ...mapGetters({
        }),
    },
    created: function () {
        console.log("ProfileView created");
    },
}
</script>

<style scoped></style>
