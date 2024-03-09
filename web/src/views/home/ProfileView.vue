<template>
    <div>
        <h2>Profile</h2>
    </div>
</template>

<script>
export default {
    props: {
        userInfo: {
            type: Object,
            default: function () {
                return {};
            }
        }
    },
    created: function () {
        console.log("ProfileView created");
    },
}
</script>

<style scoped></style>
