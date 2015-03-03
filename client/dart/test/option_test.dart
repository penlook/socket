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

library option.test;

import 'package:unittest/unittest.dart';
import 'package:unittest/html_config.dart';
import 'package:socket/socket.dart';

void main() {

    useHtmlConfiguration();

    test("option constructor", () {

        // Test default constructor
        var option = new Option();
        expect("GET", option.Method);
        expect("", option.Url);
        expect("{}", option.Data);
        expect(60, option.Timeout);
        expect(true, option.Async);

        // Test custom constructor
        option = new Option(
                     method: "POST",
                     url:"/abc",
                     data: "{key:value}",
                     timeout: 20,
                     async: false
                 );

        expect("POST", option.Method);
        expect("/abc", option.Url);
        expect("{key:value}", option.Data);
        expect(20, option.Timeout);
        expect(false, option.Async);
    });

}
