import serial
import time

class ArdInterface(object):
    port = 'COM3'
    baud = 9600
    ser = serial.Serial()
    commands = []

    def __init__(self, port = 'COM3', baud=9600):
        self.port = port
        self.baud = baud
        self.ser.baudrate = baud
        self.ser.port = port

    def setPort(self, port):
        self.ser.port = port;

    def setBaud(self, baud):
        self.ser.baudrate = baud;

    def getPort(self):
        return self.port

    def getBaud(self):
        return self.baud

    def openPort(self):
        ser.open();
        return ser.is_open

    def closePort(self):
        ser.close();

    def clearCommands(self):
        self.commands = []

    def addCommand(self,comm):
        self.commands.append(comm)
        #print(self.commands)

    def step(self):
        for command in self.commands:

            if(command[0] == "d"):
                time.sleep(float(command[1:]))
            if(command[0] == "s"):
                #first byte is window second byte is channel third byte is value
                #ex: y window, red to 255 is 0 0 255
                #ex: p window, blue, to 100 is 4 2 100
                clr = int(command[3:])
                if(clr == 255):
                    clr-=1
                if(self.ser.is_open):
                    self.ser.write(int(command[1]))
                    self.ser.write(int(command[2]))
                    self.ser.write(clr)
                print(int(command[1]))
                print(int(command[2]))
                print(int(command[3:]))
                print("\n")


def genSetCommand(color, window, value):
    command = "s"
    command += str(window)
    command += str(color)
    command += str(int(value))
    return command

def main():
    try:
        f = open("ls.li",'r')
    except IOError:
        print("Could not open file")

    a = ArdInterface("COM4", 9600)

    if(f.readline() != '#jlyon1\n'):
        exit()

    for line in f:
        line = line.replace(' ', '')
        line = line.lower()
        line = line.replace('\n','')
        if(line[0] == '#'):
            continue
        if(line[:3] == 'set'):
            n = 0;
            window = 0;
            if(line[4] == 'r'):
                n = 0
            elif(line[4] == 'g'):
                n = 1
            elif(line[4] == 'b'):
                n = 2
            else:
                n = 0
            window = int(line[5])
            com = "s"
            com += str(window)
            com += str(n)
            com += str(line[5:])
            a.addCommand(com)
        if(line[:3] == 'fdi'):
            n = 0;
            window = 0;
            window = int(line[3])
            time = float(line[4])
            r = float(line[5:8])
            g = float(line[8:11])
            b = float(line[11:14])
            ri = r/500;
            gi = g/500;
            bi = b/500;
            rc = 0;
            gc = 0;
            bc = 0;
            for i in range(500):
                a.addCommand(genSetCommand(0,window,rc))
                a.addCommand(genSetCommand(1,window,gc))
                a.addCommand(genSetCommand(2,window,bc))
                rc += ri;
                gc += gi;
                bc += bi;
                a.addCommand(("d" + str(time/500)))
            for i in range(500):
                a.addCommand(genSetCommand(0,window,rc))
                a.addCommand(genSetCommand(1,window,gc))
                a.addCommand(genSetCommand(2,window,bc))
                rc -= ri;
                gc -= gi;
                bc -= bi;
                a.addCommand(("d" + str(time/500)))


        if(line[:3] == 'fd0'):
            n = 0;
            window = 0;
            window = int(line[3])
            time = float(line[4])
            r = float(line[5:8])
            g = float(line[8:11])
            b = float(line[11:14])
            ri = 255/500;
            gi = 255/500;
            bi = 255/500;
            rc = 255;
            gc = 255;
            bc = 255;
            for i in range(500):
                a.addCommand(genSetCommand(0,window,rc))
                a.addCommand(genSetCommand(1,window,gc))
                a.addCommand(genSetCommand(2,window,bc))
                rc -= ri;
                gc -= gi;
                bc -= bi;
                a.addCommand(("d" + str(time/500)))

            print(time)

    a.step()
if __name__ == "__main__":
    main()
