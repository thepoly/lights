/* 16 Channel LED Driver Board for /The Polytechnic/.
*
* Copyright 2014 Ethan Spitz
* 
* This program is free software: you can redistribute it and/or modify
* it under the terms of the GNU General Public License as published by
* the Free Software Foundation, either version 3 of the License, or
* (at your option) any later version.
*
* This program is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
* GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License
* along with this program.  If not, see <http://www.gnu.org/licenses/>.
*
* Connections:
* # Potentiometer: A0
* # Button: D2
* # Seven-Segment Display Driver
*    # A: A1 //LSB
*    # B: 6
*    # C: 7
*    # D: A2 //MSB
* # 
*
*/

#include <Wire.h>

// definition of registers on the pwm chip
#define LED0_ON_L 0x6
#define LED0_ON_H 0x7
#define LED0_OFF_L 0x8
#define LED0_OFF_H 0x9

#define PWM_ADDR 0x40 //address of the PWM chip
int r;
int g;
int v = 0;
int b;

void setup()
{
r = 255;
g = 255;
b = 255;
// configure serial and i2c communication
Serial.begin(9600);
Wire.begin();

// configure PWM chip
setConfiguration();
  //Red
  setPWM(0,4095);
  setPWM(1,4095);
  setPWM(2,4095);
  setPWM(3,4095);
  setPWM(4,4095);
  //Green
  setPWM(10,4095);
  setPWM(11,4095); 
  setPWM(12,4095);
  setPWM(13,4095);
  setPWM(14,4095);
  //Blue
  setPWM(5,4095);
  setPWM(6,4095);
  setPWM(7,4095);
  setPWM(8,4095);
  setPWM(9,4095);
}

void serialEvent(){
  int vold = 0;
  while(v != -1){
    vold = v;
    v = (int)Serial.read();
  }
  v = vold;
}

void loop()
{  
  switch(v){
    case 0:
    r = 255;
    g = 255;
    b = 255;
    break;
    case 1:
    r = 255;
    g = 0;
    b = 0;
    break;
    case 2:
    r = 0;
    g = 255;
    b = 0;
    break;
    case 3:
    r = 0;
    g = 0;
    b = 255;
    break;
  }

    for(int i = 0; i < 4095/255; i += 1){
      setPWM(0,(int)(i*r));
      setPWM(1,(int)(i*r));
      setPWM(2,(int)(i*r));
      setPWM(3,(int)(i*r));
      setPWM(4,(int)(i*r));
      //Green
      setPWM(10,i*g);
      setPWM(11,i*g); 
      setPWM(12,i*g);
      setPWM(13,i*g);
      setPWM(14,i*g);
      //Blue
      setPWM(5,i*b);
      setPWM(6,i*b);
      setPWM(7,i*b);
      setPWM(8,i*b);
      setPWM(9,i*b);
      delay(10);
    }
    delay(1000);
    for(int i = 4094/255; i > 1; i -=1){
      setPWM(0,i*r);
      setPWM(1,i*r);
      setPWM(2,i*r);
      setPWM(3,i*r);
      setPWM(4,i*r);
      //Green
      setPWM(10,i*g);
      setPWM(11,i*g); 
      setPWM(12,i*g);
      setPWM(13,i*g);
      setPWM(14,i*g);
      //Blue
      setPWM(5,i*b);
      setPWM(6,i*b);
      setPWM(7,i*b);
      setPWM(8,i*b);
      setPWM(9,i*b);
      delay(3);
    }   
}

/*
* setPWM brightness on channel given over I2C
*/
void setPWM(int channel, uint16_t brightness)
{
  channel = (channel >= 5) ? channel += 1 : channel;
  Wire.beginTransmission(PWM_ADDR);
  Wire.write(LED0_ON_L+4*channel);
Wire.write(0x00); //turn the LED on at 0
Wire.write(0x00); //turn the LED on at 0

//turn the LED off when it hits this value (out of 4095)
Wire.write(brightness); //first four LSB
Wire.write(brightness>>8); //last four MSB
Wire.endTransmission();
}

/*
* Configure the PWM chip for easy suage and external MOSFET drive
*/
void setConfiguration()
{
  Wire.beginTransmission(PWM_ADDR);
Wire.write(0x00); //enter Mode 1 Register
Wire.write(0xa1); //enable ocsillator and auto-increment register and restart
Wire.endTransmission();
delayMicroseconds(500);//500ms delay required after reset
Wire.beginTransmission(PWM_ADDR);
Wire.write(0x01); //enter Mode 2 Register
Wire.write(0x04); //set drive mode for external MOSFETS 
Wire.endTransmission();
}
