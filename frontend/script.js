const apiBase = 'http://localhost:3000';
const historyList = [];

async function startGame() {
    const res = await fetch(`${apiBase}/start`);
    const data = await res.json();
    renderScene(data, true);
    updateLove();
}

async function sendChoice(choiceText) {
    const res = await fetch(`${apiBase}/progress`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ choice: choiceText })
    });
    const data = await res.json();
    historyList.push(`你选择了：${choiceText}`);
    renderScene(data);
    updateLove();
}

async function triggerRandomEvent() {
    const res = await fetch(`${apiBase}/random_event`);
    const data = await res.json();
    historyList.push(`事件：${data.text}`);
    renderScene(data);
}

async function updateLove() {
    const res = await fetch(`${apiBase}/state`);
    const data = await res.json();
    debugger
    document.getElementById('love').innerText = `总好感度：${data.love}`;
    // 更新每个角色的好感度
    document.getElementById('loveXiaoAi').innerText = `小艾: ${data.xiaoAi}`;
    document.getElementById('loveXiaoXue').innerText = `小雪: ${data.xiaoXue}`;
    document.getElementById('loveXiaoMei').innerText = `小美: ${data.xiaoMei}`;
}

function renderScene(scene, isStart = false) {
    if (scene.character) {
        document.getElementById('character').innerText = `[${scene.character}]`;
    } else {
        document.getElementById('character').innerText = '';
    }
    document.getElementById('text').innerText = scene.text;

    if (!isStart) {
        historyList.push(`${scene.character ? `[${scene.character}] ` : ''}${scene.text}`);
    } else {
        historyList.length = 0;
        historyList.push(`${scene.character ? `[${scene.character}] ` : ''}${scene.text}`);
    }

    renderHistory();

    const optionsDiv = document.getElementById('options');
    optionsDiv.innerHTML = '';

    if (scene.options && scene.options.length > 0) {
        scene.options.forEach(opt => {
            const btn = document.createElement('button');
            btn.innerText = opt.text;
            btn.onclick = () => sendChoice(opt.text);
            optionsDiv.appendChild(btn);
        });
    } else {
        const endText = document.createElement('div');
        endText.innerText = "游戏结束啦！🎉";
        optionsDiv.appendChild(endText);
    }
}

function renderHistory() {
    const historyUl = document.getElementById('historyList');
    historyUl.innerHTML = '';
    historyList.forEach(item => {
        const li = document.createElement('li');
        li.innerText = item;
        historyUl.appendChild(li);
    });
}

window.onload = startGame;
