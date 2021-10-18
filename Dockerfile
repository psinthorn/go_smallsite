FROM golang:1.17

# กำหนดโฟลเดอร์ที่เก็บไฟล์งานบน docker
WORKDIR /app

# ติดตั้งไฟล์และไลบรารี่ที่เกี่ยวเนื่องกับแอพพลิเคชั่น
COPY go.mod .
COPY go.sum .
RUN go mod download

# ทำสำเนาทุกไฟล์ในโฟลเดอร์จากเครื่องคอมพิวเตอร์ไปที่ docker /app 
COPY . . 

# Air use for live reload
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
CMD ["air"]

# # this is standard command to run go on docker 
# CMD ["go", "run", "main.go"]

# Remark: 
#  - every time we edit or update Dockerfile we need to re-build docker image by running 
#  - docker-compose up --build 


