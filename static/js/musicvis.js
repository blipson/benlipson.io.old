const canvas = document.getElementById("musicvis-canvas");
const gl = canvas.getContext("webgl2");
if (!gl) {
    alert("Could not initialise webgl");
}

const vertexShaderSource = `#version 300 es
in vec2 a_position;
in vec4 a_color;
uniform vec2 u_resolution;
out vec4 v_color;

void main() {
    vec2 zeroToOne = a_position / u_resolution;
    vec2 zeroToTwo = zeroToOne * 2.0;
    vec2 clipSpace = zeroToTwo - 1.0;
    gl_Position = vec4(clipSpace * vec2(1, -1), 0, 1);
    v_color = a_color;
}
`;

const fragmentShaderSource = `#version 300 es
precision mediump float;
in vec4 v_color;
out vec4 outColor;

void main() {
  outColor = v_color;
}
`;

function setRectangle(gl, x, y, width, height) {
    const x1 = x;
    const x2 = x + width;
    const y1 = y;
    const y2 = y + height;
    gl.bufferData(gl.ARRAY_BUFFER, new Float32Array([
        x1, y1,
        x2, y1,
        x1, y2,
        x1, y2,
        x2, y2,
        x2, y1]), gl.STATIC_DRAW);
}

function setColors(gl) {
    if (document.getElementById('gradient').checked) {
        gl.bufferData(
            gl.ARRAY_BUFFER,
            new Float32Array([
                Math.random(), Math.random(), Math.random(), 1,
                Math.random(), Math.random(), Math.random(), 1,
                Math.random(), Math.random(), Math.random(), 1,
                Math.random(), Math.random(), Math.random(), 1,
                Math.random(), Math.random(), Math.random(), 1,
                Math.random(), Math.random(), Math.random(), 1,
            ]),
            gl.STATIC_DRAW);

    } else {
        const r1 = Math.random();
        const b1 = Math.random();
        const g1 = Math.random();

        const r2 = Math.random();
        const b2 = Math.random();
        const g2 = Math.random();
        gl.bufferData(
            gl.ARRAY_BUFFER,
            new Float32Array([
                r1, b1, g1, 1,
                r1, b1, g1, 1,
                r1, b1, g1, 1,
                r2, b2, g2, 1,
                r2, b2, g2, 1,
                r2, b2, g2, 1,
            ]),
            gl.STATIC_DRAW);
    }
}

const vertexShader = createShader(gl, gl.VERTEX_SHADER, vertexShaderSource);
const fragmentShader = createShader(gl, gl.FRAGMENT_SHADER, fragmentShaderSource);
const program = createProgram(gl, vertexShader, fragmentShader);

const positionAttributeLocation = gl.getAttribLocation(program, "a_position");
const colorLocation = gl.getAttribLocation(program, "a_color");
const resolutionUniformLocation = gl.getUniformLocation(program, "u_resolution");

const vao = gl.createVertexArray();
gl.bindVertexArray(vao);

resize(gl.canvas);
gl.viewport(0, 0, gl.canvas.width, gl.canvas.height);
gl.clearColor(0, 0, 0, 0);
gl.clear(gl.COLOR_BUFFER_BIT);
gl.useProgram(program);
gl.uniform2f(resolutionUniformLocation, gl.canvas.width, gl.canvas.height);

function drawRectangles() {
    for (let i = 0; i < 50; ++i) {
        const positionBuffer = gl.createBuffer();
        gl.bindBuffer(gl.ARRAY_BUFFER, positionBuffer);

        setRectangle(gl, randomInt(canvas.width-200), randomInt(canvas.height-200), randomInt(300), randomInt(300));

        gl.enableVertexAttribArray(positionAttributeLocation);
        gl.enableVertexAttribArray(colorLocation);
        const size = 2;
        const type = gl.FLOAT;
        const normalize = false;
        const stride = 0;
        const vertexAttribOffset = 0;
        gl.vertexAttribPointer(positionAttributeLocation, size, type, normalize, stride, vertexAttribOffset);

        const colorBuffer = gl.createBuffer();
        gl.bindBuffer(gl.ARRAY_BUFFER, colorBuffer);

        setColors(gl);

        const colorSize = 4;
        const colorType = gl.FLOAT;
        const colorNormalize = false;
        const colorStride = 0;
        const colorOffset = 0;
        gl.vertexAttribPointer(colorLocation, colorSize, colorType, colorNormalize, colorStride, colorOffset);

        const primitiveType = gl.TRIANGLES;
        const glslOffset = 0;
        const count = 6;
        gl.drawArrays(primitiveType, glslOffset, count);
    }
}
