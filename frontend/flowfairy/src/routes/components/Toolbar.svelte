<script>
    import { LoadFlow, Ping } from "../../stores/flowstore";
    let files
</script>
<div class="toolbar-container">
    <div class="logo-container">
        <p class="toolbar-text01">FlowFairy</p>

    </div>
    <div class="util-container">
        <input class="toolbar-util" type="file" multiple bind:files on:change={() => LoadFlow(files)}>

        <button class="toolbar-util" on:click={Ping}>
            S
        </button>

        <button class="toolbar-util">
            T
        </button>
    </div>
</div>

<style>
     :root {
        --header-bg-color: #040a14;
     }

    .toolbar-container {
        display: flex;
        flex-direction: column;
        align-items: center;
        width: 10vw;
        height: 100vh;
        background-color: var(--header-bg-color);
        color: #ffffff;
    }

    .logo-container {
        display: flex;
        flex-direction: row;
    }

    .toolbar-text01 {
        font-size: 24px;
    }

    .util-container {
        display: flex;
        flex-direction: column;
        height: 20vh;
        justify-content: space-evenly;
        width: 75%;
        margin-top: 1vh;
        margin-left: auto;
        margin-right: auto;
    }

    .toolbar-util {
        height: 12%;
    }

    .toolbar-util:hover {
        cursor: pointer;
    }
</style>