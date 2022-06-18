<script>
    import Toolbar from "./components/Toolbar.svelte";
    import Viewer from "./components/Viewer.svelte";
</script>

<body>
    <Viewer />
    <Toolbar />

</body>

<style>
    :root {
        --bg-color: #f2f4f8;
        --dark-text-color: #101010;
        --light-text-color: #ffffff;
        --secondary-light-text-color: #a7a7a7;
    }

    body {
        padding: 0;
        margin: 0;
        width: 100vw;
        height: 100vh;
        box-sizing: border-box;
        font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
        background-color: var(--bg-color);
    }
</style>