package com.cborum;

import org.junit.Test;

import static junit.framework.TestCase.assertEquals;

public class TestMyArrayListWithBugs {

    @Test
    public void testAddObj() {
        MyArrayListWithBugs mList = new MyArrayListWithBugs();
        mList.add("a");
        assertEquals(1, mList.size());
    }

    @Test
    public void testAddIndexObj1() {
        MyArrayListWithBugs mList = new MyArrayListWithBugs();
        mList.add("a");
        mList.add("b");
        mList.add(0, "c");
        assertEquals(3, mList.size());
    }

    @Test
    public void testAddIndexObj2() {
        MyArrayListWithBugs mList = new MyArrayListWithBugs();
        mList.add("a");
        mList.add("b");
        mList.add(1, "c");
        assertEquals(3, mList.size());
        String obj = (String) mList.remove(1);
        assertEquals("c", obj);
        assertEquals(2, mList.size());
    }

    @Test
    public void testGetObj() {
        MyArrayListWithBugs mList = new MyArrayListWithBugs();
        mList.add("a");
        mList.add("b");
        Object obj = mList.get(0);
        System.out.println(obj);
    }

    @Test (expected = IndexOutOfBoundsException.class)
    public void testGetOutOfRange() {
        MyArrayListWithBugs mList = new MyArrayListWithBugs();
        mList.add("a");
        mList.add("b");
        Object obj = mList.get(2);
        System.out.println(obj);

        /*try {
            Object obj = mList.get(2);
            System.out.println(obj);
        } catch (Exception e) {
            assertEquals(IndexOutOfBoundsException.class, e.getClass());
        }*/
    }

    @Test
    public void testGetLastObj() {
        MyArrayListWithBugs mList = new MyArrayListWithBugs();
        mList.add("a");
        mList.add("b");
        mList.add("d");
        String obj = (String) mList.get(2);
        assertEquals("d", obj);
    }

    @Test
    public void testRemoveObj() {
        MyArrayListWithBugs mList = new MyArrayListWithBugs();
        mList.add("a");
        mList.add("b");
        Object obj = mList.remove(0);
        System.out.println(obj);
        assertEquals("a", (String) obj);
    }

    @Test
    public void testIncreaseLiseSize() {
        MyArrayListWithBugs mList = new MyArrayListWithBugs();
        mList.add("a");
        mList.add("b");
        mList.add("c");
        mList.add("d");
        mList.add("1");
        mList.add("2");
        mList.add("3");
        mList.add("4");
        mList.add("5");
        mList.add("6");
        mList.add("7");
        assertEquals(11, mList.size());
    }
}
