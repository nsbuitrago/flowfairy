
<div class="toolbar-container">
    <div class="app-logo-container">
        <h2>FlowFairy</h2>
    </div>
    <div class="utility-container">
        <label for="file-dialog">Load Data</label>
        <input type="file" multiple class="utility-button" id="file-dialog">
        <button class="utility-button">Save Data</button>
    </div>
</div>

<style>
    .toolbar-container {
        display: flexbox;
        position: fixed;
        flex-direction: column;
        height: 100vh;
        width: 12rem;
        top: 0;
        left: 0;
        border-top-right-radius: 10px;
        background-color: black;
        font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
    }

    .utility-container {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: space-evenly;
        height: 20vh;
    }

    .app-logo-container {
        display: flex;
        justify-content: center;
        color: white;
    }

    .utility-button {
        width: 50%;
        padding: 2%;
    }

    .utility-button:hover {
        background-color: white;
        cursor: pointer;
        transition: ease-in-out 300m;
    }
</style>