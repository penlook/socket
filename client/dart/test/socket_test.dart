/**
 * Penlook Project
 *
 * Copyright (c) 2015 Penlook Development Team
 *
 * --------------------------------------------------------------------
 *
 * This program is free software: you can redistribute it and/or
 * modify it under the terms of the GNU Affero General Public License
 * as published by the Free Software Foundation, either version 3
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public
 * License along with this program.
 * If not, see <http://www.gnu.org/licenses/>.
 *
 * --------------------------------------------------------------------
 *
 * Authors:
 *     Loi Nguyen       <loint@penlook.com>
 *     Nam Vo           <namvh@penlook.com>
 */

library socket.test;

import 'package:unittest/unittest.dart';
import 'package:unittest/html_config.dart';
import 'package:socket/socket.dart';

void main() {

    useHtmlConfiguration();

    test("socket constructor", () {
          
        var socket = new Socket();
        
        expect("localhost", socket.Host);
        expect(80, socket.Port);
        
        socket = new Socket(host:"192.168.2.1");
        expect("192.168.2.1", socket.Host);
        expect(80, socket.Port);
        
        socket = new Socket(host:"127.0.0.1", port: 3000);
        expect("127.0.0.1", socket.Host);
        expect(3000, socket.Port);
       
    });
    
    test("synchronous request",() {
        
        var socket = new Socket();  
        expect("test", "test");        
        
    }); 
    
    test("asynchronous request",() {
        
        var socket = new Socket();  
        expect("test", "test");        
        
    }); 
    
    
    test("asynchronous request",() {
            
        var socket = new Socket();  
        expect("test", "test");        
        
    }); 
    
    
        
    

}