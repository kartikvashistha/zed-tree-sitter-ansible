import XCTest
import SwiftTreeSitter
import TreeSitterAnsible

final class TreeSitterAnsibleTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_ansible())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Ansible grammar")
    }
}
